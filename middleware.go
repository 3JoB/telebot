package telebot

type Middlewares struct {
	Before []MiddlewareFunc
	After  []MiddlewareFunc
	Final  []MiddlewareFunc
}

// MiddlewareFunc represents a middleware processing function,
// which get called before the endpoint group or specific handler.
type MiddlewareFunc func(HandlerFunc) HandlerFunc

func applyMiddleware(h HandlerFunc, m ...MiddlewareFunc) HandlerFunc {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

// Group is a separated group of handlers, united by the general middleware.
type Group struct {
	b          *Bot
	middleware []MiddlewareFunc
}

// Use adds middleware to the chain.
func (g *Group) Use(middleware ...MiddlewareFunc) {
	g.middleware = append(g.middleware, middleware...)
}

// Handle adds endpoint handler to the bot, combining group's middleware
// with the optional given middleware.
func (g *Group) Handle(endpoint any, h HandlerFunc, m ...MiddlewareFunc) {
	mw := m
	if len(g.middleware) > 0 {
		mw = make([]MiddlewareFunc, 0, len(g.middleware)+len(m))
		mw = append(mw, g.middleware...)
		mw = append(mw, m...)
	}
	g.b.Handle(endpoint, h, mw...)
}
