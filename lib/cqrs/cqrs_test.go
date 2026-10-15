package cqrs

import (
	"context"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

	"modular-monolith/lib/config"
	"modular-monolith/lib/events"
)

// ---- test messages (exported: StructName requires reflection-safe names) ----

type PingCmd struct{ Msg string }

type Ponged struct{ Msg string }

// ---- handlers ----

type pingHandler struct {
	done chan string
	bus  EventBus
}

func (h *pingHandler) HandlerName() string { return "PingHandler" }
func (h *pingHandler) NewCommand() any     { return &PingCmd{} }
func (h *pingHandler) Handle(ctx context.Context, c any) error {
	cmd := c.(*PingCmd)
	h.done <- cmd.Msg
	return h.bus.Publish(ctx, &Ponged{Msg: cmd.Msg})
}

type pongHandler struct{ done chan string }

func (h *pongHandler) HandlerName() string { return "PongHandler" }
func (h *pongHandler) NewEvent() any       { return &Ponged{} }
func (h *pongHandler) Handle(_ context.Context, e any) error {
	h.done <- e.(*Ponged).Msg
	return nil
}

func setup(t *testing.T) (CommandBus, EventBus, *pingHandler, *pongHandler) {
	t.Helper()
	logger := watermill.NopLogger{}
	ps := &events.PubSub{}
	gc := gochannel.NewGoChannel(gochannel.Config{}, logger)
	ps.Publisher, ps.Subscriber = gc, gc

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = router.Close() })

	cmdBusIface, err := NewCommandBus(ps, enabledCfg(), logger)
	if err != nil {
		t.Fatal(err)
	}
	evtBusIface, err := NewEventBus(ps, enabledCfg(), logger)
	if err != nil {
		t.Fatal(err)
	}

	ping := &pingHandler{done: make(chan string, 1), bus: evtBusIface}
	pong := &pongHandler{done: make(chan string, 1)}

	reg := NewRegistry(registryParams{
		Commands: []CommandHandler{ping},
		Events:   []EventHandler{pong},
		Cfg:      enabledCfg(),
		PubSub:   ps,
		Logger:   logger,
	})
	if err := reg.Attach(router); err != nil {
		t.Fatal(err)
	}

	go func() {
		_ = router.Run(context.Background())
	}()
	waitRunning(t, router)

	return cmdBusIface, evtBusIface, ping, pong
}

func enabledCfg() *config.Config {
	cfg := &config.Config{}
	cfg.Watermill.CommandBus.Enabled = true
	cfg.Watermill.EventBus.Enabled = true
	return cfg
}

func waitRunning(t *testing.T, r *message.Router) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if r.IsRunning() {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("router did not start")
}

func waitFor(t *testing.T, ch chan string, want string) {
	t.Helper()
	select {
	case got := <-ch:
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("timeout waiting for message")
	}
}

func TestCommandAndEventRoundTrip(t *testing.T) {
	cmdBus, _, ping, pong := setup(t)

	if err := cmdBus.Send(context.Background(), &PingCmd{Msg: "hello"}); err != nil {
		t.Fatalf("Send: %v", err)
	}
	// command handler runs, then publishes Ponged; event handler receives it
	waitFor(t, ping.done, "hello")
	waitFor(t, pong.done, "hello")
}
