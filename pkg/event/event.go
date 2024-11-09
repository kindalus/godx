package event

import (
	"errors"
	"time"

	"github.com/kindalus/godx/pkg/nanoid"
)

const (
	HeaderEventID       = "EventID"
	HeaderEventName     = "EventName"
	HeaderIssuedAt      = "IssuedAt"
	HeaderAggregateType = "AggregateType"
	HeaderAggregateID   = "AggregateID"
	HeaderCommandID     = "CommandID"
)

var ErrEventNotFound = errors.New("event not found")

type Event struct {
	headers map[string]string
	payload interface{}
}

func New(name string, opts ...Option) Event {
	id := nanoid.New().String()
	myHeaders := make(map[string]string)
	myHeaders[HeaderEventName] = name
	myHeaders[HeaderEventID] = id
	myHeaders[HeaderIssuedAt] = time.Now().Format(time.RFC3339)

	e := Event{headers: myHeaders, payload: nil}

	for _, opt := range opts {
		e = opt(e)
	}

	return e
}

type Option func(Event) Event

func WithPayload(b interface{}) Option {
	return func(e Event) Event {
		e.payload = b
		return e
	}
}

type HeaderEntry [2]string

func WithHeader(k string, v string) Option {
	return func(e Event) Event {
		e.headers[k] = v
		return e
	}
}

func WithHeaders(h ...HeaderEntry) Option {
	return func(evt Event) Event {
		for _, e := range h {
			evt.headers[e[0]] = e[1]
		}
		return evt
	}
}

func WithOptions(opts ...Option) Option {
	return func(evt Event) Event {
		for _, opt := range opts {
			evt = opt(evt)
		}
		return evt
	}
}

// DecorateEvent applies additional options to an existing event, allowing it to be modified or "decorated" in-place.
// This is useful for cases where you want to add headers, modify the payload, or update the context of an existing event.
func Decorate(e Event, opts ...Option) Event {
	empty := New(e.Name())
	cloned := Clone(e)

	for k, v := range empty.headers {
		cloned.headers[k] = v
	}

	for _, opt := range opts {
		cloned = opt(cloned)
	}

	return cloned
}

func Clone(e Event) Event {
	entries := make([]HeaderEntry, 0, len(e.headers))
	for k, v := range e.headers {
		entries = append(entries, HeaderEntry{k, v})
	}

	return New(
		e.Name(),
		WithHeaders(entries...),
		WithPayload(e.Payload()),
	)
}

func (e Event) ID() string {
	return e.headers[HeaderEventID]
}

func (e Event) Name() string {
	return e.headers[HeaderEventName]
}

func (e Event) IssuedAt() string {
	t := e.headers[HeaderIssuedAt]
	return t
}

func (e Event) Header(key string) string {
	return e.headers[key]
}

func (e Event) Payload() interface{} {
	return e.payload
}
