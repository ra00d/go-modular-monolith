package events

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"

	"modular-monolith/lib/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// testConfig builds a config pointing at the postgres transport.
func postgresCfg(dsn string) *config.Config {
	cfg := &config.Config{}
	cfg.Watermill.Transport = TransportPostgres
	cfg.Watermill.Postgres.Schema = "watermill_test"
	// Reuse DatabaseConfig parsing via koanf-free direct construction.
	cfg.Database.Driver = "postgres"
	cfg.Database.Host = os.Getenv("TEST_PG_HOST")
	cfg.Database.Port = 5433
	cfg.Database.Database = os.Getenv("TEST_PG_DB")
	cfg.Database.Username = os.Getenv("TEST_PG_USER")
	cfg.Database.Password = os.Getenv("TEST_PG_PASSWORD")
	if cfg.Database.Host == "" {
		cfg.Database.Host = "localhost"
		cfg.Database.Database = "erp"
		cfg.Database.Username = "postgres"
		cfg.Database.Password = "secret"
	}
	cfg.Database.SSLMode = "disable"
	return cfg
}

// TestPostgresRoundTrip verifies the durable transport: publish, consume,
// ack. Skipped when Postgres is unreachable so CI without infra stays green.
func TestPostgresRoundTrip(t *testing.T) {
	cfg := postgresCfg("")
	db, err := openDB(cfg)
	if err != nil {
		t.Skipf("postgres unreachable: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	logger := watermill.NopLogger{}
	ps, err := NewPubSub(cfg, db, logger)
	if err != nil {
		t.Fatalf("NewPubSub: %v", err)
	}
	t.Cleanup(func() { _ = ps.Close() })

	const topic = "lib_events_roundtrip"
	received := make(chan *message.Message, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		messages, err := ps.Subscriber.Subscribe(ctx, topic)
		if err != nil {
			return
		}
		for msg := range messages {
			received <- msg
			msg.Ack()
		}
	}()

	time.Sleep(200 * time.Millisecond) // let the subscriber poll loop attach

	payload := []byte(`{"hello":"world"}`)
	if err := ps.Publisher.Publish(topic, message.NewMessage(watermill.NewUUID(), payload)); err != nil {
		t.Fatalf("Publish: %v", err)
	}

	select {
	case msg := <-received:
		if string(msg.Payload) != string(payload) {
			t.Fatalf("payload = %s", msg.Payload)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("timeout waiting for published message")
	}
}

// openDB connects using the same DSN builder as the application.
func openDB(cfg *config.Config) (*sqlx.DB, error) {
	dsn, err := cfg.Database.DSN()
	if err != nil {
		return nil, err
	}
	return sqlx.Connect("postgres", dsn)
}
