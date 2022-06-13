package common

import (
	"context"
)

type Snapshot struct {
	ID      string        `json:"id"`
	Type    AggregateType `json:"type"`
	State   []byte        `json:"state"`
	Version uint64        `json:"version"`
}

type ILoad interface {
	Load(events []Event) error
}

type IApply interface {
	Apply(evt Event) error
}

type IAggregateRoot interface {
	GetUncommittedEvents() []Event
	GetID() string
	SetID(id string) *AggregateBase
	GetVersion() int64
	ClearUncommittedEvents()
	ToSnapshot()
	SetType(aggregateType AggregateType)
	GetType() AggregateType
	SetAppliedEvents(events []Event)
	GetAppliedEvents() []Event
	RaiseEvent(event Event) error
	String() string
}

type IAggregate interface {
	IAggregateRoot
	When(event Event) error
}

type IAggregateStore interface {
	// Load loads the most recent version of an aggregate to provided  into params aggregate with a type and id.
	Load(ctx context.Context, aggregate IAggregate) error
	// Save saves the uncommitted events for an aggregate.
	Save(ctx context.Context, aggregate IAggregate) error
	// Exists check aggregate exists by id.
	Exists(ctx context.Context, streamID string) error
	// IEventStore
	// ISnapshotStore
}

// EventStore is an interface for an event sourcing event store.
type IEventStore interface {
	// SaveEvents appends all events in the event stream to the store.
	SaveEvents(ctx context.Context, streamID string, events []Event) error

	// LoadEvents loads all events for the aggregate id from the store.
	LoadEvents(ctx context.Context, streamID string) ([]Event, error)
}

// ISnapshotStore is an interface for an event sourcing snapshot store.
type ISnapshotStore interface {
	// SaveSnapshot save aggregate snapshot.
	SaveSnapshot(ctx context.Context, aggregate IAggregate) error

	// GetSnapshot load aggregate snapshot.
	GetSnapshot(ctx context.Context, id string) (*Snapshot, error)
}
