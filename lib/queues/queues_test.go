package queues

import (
	"context"
	stderrors "errors"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

	"modular-monolith/lib/config"
	"modular-monolith/lib/events"
)

type emailPayload struct{ To string }

func TestQueueRoundTrip(t *testing.T) {
	logger := watermill.NopLogger{}
	gc := gochannel.NewGoChannel(gochannel.Config{}, logger)
	ps := &events.PubSub{Publisher: gc, Subscriber: gc}

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = router.Close() })

	cfg := &config.Config{}
	cfg.Queues.Enabled = true
	cfg.Queues.RetryCount = 0
	q := NewQueue(ps, cfg)

	received := make(chan string, 1)
	if err := AddWorker(router, q, "send_email", func(ctx context.Context, p *emailPayload) error {
		received <- p.To
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	go func() { _ = router.Run(context.Background()) }()
	waitRunning(t, router)

	if err := q.Enqueue(context.Background(), "send_email", &emailPayload{To: "dev@qistas.io"}); err != nil {
		t.Fatalf("Enqueue: %v", err)
	}

	select {
	case got := <-received:
		if got != "dev@qistas.io" {
			t.Fatalf("got %q", got)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("timeout waiting for job")
	}
}

func TestQueueDisabled(t *testing.T) {
	cfg := &config.Config{} // Queues.Enabled defaults false
	q := NewQueue(&events.PubSub{}, cfg)

	if err := q.Enqueue(context.Background(), "job", &emailPayload{}); !stderrors.Is(err, ErrQueuesDisabled) {
		t.Errorf("Enqueue err = %v, want ErrQueuesDisabled", err)
	}
	// AddWorker checks enabled before touching the router; a nil router is
	// safe because the disabled check short-circuits first.
	var nilRouter *message.Router
	if err := AddWorker(nilRouter, q, "job", func(context.Context, *emailPayload) error { return nil }); !stderrors.Is(err, ErrQueuesDisabled) {
		t.Errorf("AddWorker err = %v, want ErrQueuesDisabled", err)
	}
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
