package events

import "context"

// Subscriber is abstract subscriber for the event
type Subscriber interface {
	HandleEvent(ctx context.Context, event Event)
}

type SubscriberFunc func(ctx context.Context, event Event)

func (fn SubscriberFunc) HandleEvent(ctx context.Context, event Event) {
	fn(ctx, event)
}
