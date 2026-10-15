// Package cqrs wires Watermill's CQRS component into the application:
//
//   - CommandBus: Send(ctx, cmd) — consumed by exactly one handler
//   - EventBus:   Publish(ctx, evt) — broadcast notifications
//
// Handlers implement the watermill cqrs.Handler interfaces (build them with
// [CommandHandler] / [EventHandler] constructors) and are collected into a
// Registry through fx value groups:
//
//	fx.Provide(fx.Annotate(
//	    NewAssignPlanHandler,
//	    fx.As(new(cqrslib.CommandHandler)),
//	    fx.ResultTags(`group:"commands"`),
//	))
//
// The registry attaches processors to the process-wide router during app
// construction, before the router starts (events.Module lifecycle hook).
package cqrs

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/fx"

	"modular-monolith/lib/config"
	"modular-monolith/lib/events"
)

// Handler aliases so modules annotate against this package without importing
// watermill's cqrs component directly.
type (
	CommandHandler = cqrs.CommandHandler
	EventHandler   = cqrs.EventHandler
)

var marshaler = &cqrs.JSONMarshaler{GenerateName: cqrs.StructName}

func commandTopic(name string) string {
	return events.CommandsTopicPrefix + name
}

func eventTopic(name string) string {
	return events.EventsTopicPrefix + name
}

// CommandBus dispatches commands to their single handler.
type CommandBus interface {
	Send(ctx context.Context, cmd any) error
}

// EventBus broadcasts domain events to every subscribed handler.
type EventBus interface {
	Publish(ctx context.Context, evt any) error
}

// ---------- implementations ----------

type commandBus struct{ inner *cqrs.CommandBus }

func (b commandBus) Send(ctx context.Context, cmd any) error {
	return (*b.inner).Send(ctx, cmd)
}

type nopCommandBus struct{}

func (nopCommandBus) Send(context.Context, any) error {
	return nil // disabled by config; sends are dropped
}

type eventBus struct{ inner *cqrs.EventBus }

func (b eventBus) Publish(ctx context.Context, evt any) error {
	return (*b.inner).Publish(ctx, evt)
}

type nopEventBus struct{}

func (nopEventBus) Publish(context.Context, any) error {
	return nil // disabled by config; publishes are dropped
}

// ---------- constructors (fx providers) ----------

// NewCommandBus builds the command bus over the shared publisher.
// Disabled in config -> a no-op bus keeps module wiring stable.
func NewCommandBus(ps *events.PubSub, cfg *config.Config, logger watermill.LoggerAdapter) (CommandBus, error) {
	if !cfg.Watermill.CommandBus.Enabled {
		logger.Info("cqrs: command bus disabled by config", nil)
		return nopCommandBus{}, nil
	}
	bus, err := cqrs.NewCommandBusWithConfig(ps.Publisher, cqrs.CommandBusConfig{
		GeneratePublishTopic: func(params cqrs.CommandBusGeneratePublishTopicParams) (string, error) {
			return commandTopic(params.CommandName), nil
		},
		Marshaler: marshaler,
		Logger:    logger,
	})
	if err != nil {
		return nil, fmt.Errorf("cqrs: command bus: %w", err)
	}
	return commandBus{inner: bus}, nil
}

// NewEventBus builds the event bus over the shared publisher.
func NewEventBus(ps *events.PubSub, cfg *config.Config, logger watermill.LoggerAdapter) (EventBus, error) {
	if !cfg.Watermill.EventBus.Enabled {
		logger.Info("cqrs: event bus disabled by config", nil)
		return nopEventBus{}, nil
	}
	bus, err := cqrs.NewEventBusWithConfig(ps.Publisher, cqrs.EventBusConfig{
		GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {
			return eventTopic(params.EventName), nil
		},
		Marshaler: marshaler,
		Logger:    logger,
	})
	if err != nil {
		return nil, fmt.Errorf("cqrs: event bus: %w", err)
	}
	return eventBus{inner: bus}, nil
}

// ---------- registry ----------

type registryParams struct {
	fx.In

	Commands []CommandHandler `group:"commands"`
	Events   []EventHandler   `group:"events"`

	Cfg    *config.Config
	PubSub *events.PubSub
	Logger watermill.LoggerAdapter
}

// Registry aggregates every handler provided under the "commands" /
// "events" fx value groups and attaches their processors to the router.
type Registry struct {
	params registryParams
}

// NewRegistry builds the registry; fx injects all grouped handlers here,
// so it is complete before Attach runs.
func NewRegistry(p registryParams) *Registry {
	return &Registry{params: p}
}

// Handlers reports how many handlers are registered per kind.
func (r *Registry) Handlers() (commands, eventsCount int) {
	return len(r.params.Commands), len(r.params.Events)
}

// Attach registers processors onto the router. Called during app
// construction; the router starts afterwards via the events module hook.
func (r *Registry) Attach(router *message.Router) error {
	cfg := r.params.Cfg
	ps := r.params.PubSub
	log := r.params.Logger

	if cfg.Watermill.CommandBus.Enabled && len(r.params.Commands) > 0 {
		processor, err := cqrs.NewCommandProcessorWithConfig(router, cqrs.CommandProcessorConfig{
			GenerateSubscribeTopic: func(params cqrs.CommandProcessorGenerateSubscribeTopicParams) (string, error) {
				return commandTopic(params.CommandName), nil
			},
			SubscriberConstructor: func(cqrs.CommandProcessorSubscriberConstructorParams) (message.Subscriber, error) {
				return ps.Subscriber, nil
			},
			Marshaler: marshaler,
			Logger:    log,
		})
		if err != nil {
			return fmt.Errorf("cqrs: command processor: %w", err)
		}
		if err := processor.AddHandlers(r.params.Commands...); err != nil {
			return fmt.Errorf("cqrs: add command handlers: %w", err)
		}
	}

	if cfg.Watermill.EventBus.Enabled && len(r.params.Events) > 0 {
		processor, err := cqrs.NewEventProcessorWithConfig(router, cqrs.EventProcessorConfig{
			GenerateSubscribeTopic: func(params cqrs.EventProcessorGenerateSubscribeTopicParams) (string, error) {
				return eventTopic(params.EventName), nil
			},
			SubscriberConstructor: func(cqrs.EventProcessorSubscriberConstructorParams) (message.Subscriber, error) {
				return ps.Subscriber, nil
			},
			Marshaler: marshaler,
			Logger:    log,
		})
		if err != nil {
			return fmt.Errorf("cqrs: event processor: %w", err)
		}
		if err := processor.AddHandlers(r.params.Events...); err != nil {
			return fmt.Errorf("cqrs: add event handlers: %w", err)
		}
	}
	return nil
}

// Module provides the buses and the handler registry.
var Module = fx.Module("cqrs",
	fx.Provide(NewCommandBus),
	fx.Provide(NewEventBus),
	fx.Provide(NewRegistry),
	fx.Invoke(func(reg *Registry, router *message.Router) error {
		return reg.Attach(router)
	}),
)
