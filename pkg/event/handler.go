package event

type Handler interface {
	Handle(event Event)
}

type HandlerFunc func(Event)

func (f HandlerFunc) Handle(event Event) {
	f(event)
}

type Middleware func(Handler) Handler

func Chain(s Handler, m ...Middleware) Handler {
	for i := len(m) - 1; i >= 0; i-- {
		s = m[i](s)
	}
	return s
}
