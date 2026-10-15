// Package queues implements background job execution over Watermill:
//
//	q := queues.New(...)                       // fx-provided
//	queues.AddWorker(router, q, "send_email", func(ctx context.Context, p *EmailPayload) error { ... })
//	q.Enqueue(ctx, "send_email", &EmailPayload{...})
//
// Jobs publish to the "queue.<name>" topic and are consumed by their worker
// through the shared router; failures retry per QueuesConfig (RetryCount /
// RetryDelay middleware on the router) and land on the poison topic after
// the last attempt. When queues are disabled in config, Enqueue and
// AddWorker return [ErrQueuesDisabled] — callers decide how to degrade.
package queues

import (
	"context"

	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"

	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/fx"

	"modular-monolith/lib/config"
	"modular-monolith/lib/events"
)

// DisabledError is returned by Queue operations when queues.enabled is false.
var ErrQueuesDisabled = fmt.Errorf("queues: disabled by config")

// Queue enqueues background jobs onto the messaging transport.
type Queue struct {
	pub message.Publisher
	sub message.Subscriber

	enabled bool
}

// NewQueue builds the job queue over the shared transport.
func NewQueue(ps *events.PubSub, cfg *config.Config) *Queue {
	return &Queue{pub: ps.Publisher, sub: ps.Subscriber, enabled: cfg.Queues.Enabled}
}

// Module provides the job queue.
var Module = fx.Module("queues", fx.Provide(NewQueue))

// Topic returns the queue topic for a job name.
func Topic(job string) string {
	return events.QueueTopicPrefix + job
}

// Enqueue publishes payload for asynchronous processing by the job's worker.
func (q *Queue) Enqueue(ctx context.Context, job string, payload any) error {
	if !q.enabled {
		return ErrQueuesDisabled
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("queues: marshal %s: %w", job, err)
	}
	msg := message.NewMessage(watermill.NewUUID(), body)
	if err := q.pub.Publish(Topic(job), msg); err != nil {
		return fmt.Errorf("queues: enqueue %s: %w", job, err)
	}
	return nil
}

// AddWorker subscribes fn to the job's topic on the shared router. The type
// parameter is the decoded payload type. Must run before the router starts
// (i.e. during app construction, not from lifecycle hooks).
func AddWorker[P any](router *message.Router, q *Queue, name string, fn func(ctx context.Context, p *P) error) error {
	if !q.enabled {
		return ErrQueuesDisabled
	}
	router.AddNoPublisherHandler(
		name,
		Topic(name),
		q.sub,
		func(msg *message.Message) error {
			var p P
			if len(msg.Payload) > 0 {
				if err := json.Unmarshal(msg.Payload, &p); err != nil {
					return fmt.Errorf("queues: decode %s: %w", name, err)
				}
			}
			return fn(msg.Context(), &p)
		},
	)
	return nil
}
