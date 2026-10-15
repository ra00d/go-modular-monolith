// Package events provides the shared Watermill messaging infrastructure:
// a transport (in-memory or PostgreSQL) selected by config, and the message
// router that command, event, and queue handlers run on.
//
// Transports:
//   - "memory"   gochannel pub/sub; dev/test only, messages are not durable
//   - "postgres" watermill-sql; durable, at-least-once, tables auto-created
//     in the configured schema
package events

import (
	"context"
	"fmt"
	"time"

	"modular-monolith/lib/config"
	"modular-monolith/lib/logger"

	"github.com/ThreeDotsLabs/watermill"
	sql "github.com/ThreeDotsLabs/watermill-sql/v3/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

// Transport names accepted by config.WatermillConfig.Transport.
const (
	TransportMemory   = "memory"
	TransportPostgres = "postgres"
)

// Topic prefixes keep the message namespace tidy. Modules publish domain
// events to their own topic constants built on these.
const (
	EventsTopicPrefix   = "events."
	CommandsTopicPrefix = "commands."
	QueueTopicPrefix    = "queue."
)

const poisonTopic = EventsTopicPrefix + "poison"

// PubSub bundles the transport's publisher and subscriber pair.
type PubSub struct {
	Publisher  message.Publisher
	Subscriber message.Subscriber
}

// Close releases both ends of the transport.
func (p *PubSub) Close() error {
	var firstErr error
	if err := p.Publisher.Close(); err != nil {
		firstErr = err
	}
	if err := p.Subscriber.Close(); err != nil && firstErr == nil {
		firstErr = err
	}
	return firstErr
}

// NewPubSub builds the publisher/subscriber pair for the configured
// transport. The postgres transport shares the application's *sqlx.DB pool.
func NewPubSub(cfg *config.Config, db *sqlx.DB, wmLogger watermill.LoggerAdapter) (*PubSub, error) {
	// Startup-time schema bootstrap; not request-scoped.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	switch cfg.Watermill.Transport {
	case TransportPostgres:
		schema := cfg.Watermill.Postgres.Schema
		if schema == "" {
			schema = "public"
		}
		// watermill-sql creates its tables but never the schema itself.
		if _, err := db.ExecContext(ctx, fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS %q`, schema)); err != nil {
			return nil, fmt.Errorf("events: create schema %s: %w", schema, err)
		}
		adapter := sql.DefaultPostgreSQLSchema{
			GenerateMessagesTableName: func(topic string) string {
				return fmt.Sprintf(`"%s"."watermill_%s"`, schema, topic)
			},
		}
		offsets := sql.DefaultPostgreSQLOffsetsAdapter{
			GenerateMessagesOffsetsTableName: func(topic string) string {
				return fmt.Sprintf(`"%s"."watermill_offsets_%s"`, schema, topic)
			},
		}
		publisher, err := sql.NewPublisher(db.DB,
			sql.PublisherConfig{
				SchemaAdapter:        adapter,
				AutoInitializeSchema: true,
			},
			wmLogger)
		if err != nil {
			return nil, fmt.Errorf("events: postgres publisher: %w", err)
		}
		subscriber, err := sql.NewSubscriber(db.DB,
			sql.SubscriberConfig{
				SchemaAdapter:    adapter,
				OffsetsAdapter:   offsets,
				InitializeSchema: true,
				ConsumerGroup:    "qistas",
			},
			wmLogger)
		if err != nil {
			return nil, fmt.Errorf("events: postgres subscriber: %w", err)
		}
		return &PubSub{Publisher: publisher, Subscriber: subscriber}, nil

	case TransportMemory, "":
		ps := gochannel.NewGoChannel(gochannel.Config{}, wmLogger)
		return &PubSub{Publisher: ps, Subscriber: ps}, nil

	default:
		return nil, fmt.Errorf("events: unknown transport %q", cfg.Watermill.Transport)
	}
}

// NewRouter builds the process-wide router with the standard middleware
// stack: panic recovery, correlation ids, retries driven by QueuesConfig,
// and a poison queue so undeliverable messages are parked instead of lost.
// Handlers register onto it via lib/cqrs and lib/queues before start.
func NewRouter(
	cfg *config.Config,
	ps *PubSub,
	wmLogger watermill.LoggerAdapter,
) (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{
		CloseTimeout: closeTimeout(cfg),
	}, wmLogger)
	if err != nil {
		return nil, fmt.Errorf("events: router: %w", err)
	}

	poison, err := middleware.PoisonQueue(ps.Publisher, poisonTopic)
	if err != nil {
		return nil, fmt.Errorf("events: poison queue: %w", err)
	}

	router.AddMiddleware(
		middleware.Recoverer,
		middleware.CorrelationID,
		middleware.Retry{
			MaxRetries:      cfg.Queues.RetryCount,
			InitialInterval: retryDelay(cfg),
			Logger:          wmLogger,
		}.Middleware,
		poison,
	)
	return router, nil
}

func retryDelay(cfg *config.Config) time.Duration {
	const fallback = 30 * time.Second
	raw := cfg.Queues.RetryDelay
	if raw == "" {
		return fallback
	}
	d, err := time.ParseDuration(raw)
	if err != nil || d <= 0 {
		return fallback
	}
	return d
}

func closeTimeout(cfg *config.Config) time.Duration {
	const fallback = 30 * time.Second
	raw := cfg.Watermill.Router.CloseTimeout
	if raw == "" {
		return fallback
	}
	d, err := time.ParseDuration(raw)
	if err != nil || d <= 0 {
		return fallback
	}
	return d
}

// RunRouter starts the router on a background context and stops it cleanly
// on shutdown. Intended as an fx lifecycle hook.
func RunRouter(lc fx.Lifecycle, r *message.Router) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				// Router.Run blocks until stop; a background context keeps
				// shutdown driven by OnStop's CloseTimeout, not request scopes.
				if err := r.Run(context.Background()); err != nil {
					logger.Error("events: router stopped with error:", err)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			return r.Close()
		},
	})
}

// Module wires messaging into fx.
var Module = fx.Module("events",
	fx.Provide(logger.NewWatermillLogger),
	fx.Provide(NewPubSub),
	fx.Provide(NewRouter),
	fx.Invoke(RunRouter),
)
