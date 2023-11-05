package crare

// Group is a separated group of handlers, united by the general middleware.
type Group struct {
	b          *Bot
	middleware []HandlerFunc
}

// Use adds middleware to the chain.
func (g *Group) Use(middleware ...HandlerFunc) {
	g.middleware = append(g.middleware, middleware...)
}

// Handle adds endpoint handler to the bot, combining group's middleware
// with the optional given middleware.
func (g *Group) Handle(endpoint any, h HandlerFunc, m ...HandlerFunc) {
	mw := m
	if len(g.middleware) > 0 {
		mw = make([]HandlerFunc, 0, len(g.middleware)+len(m))
		mw = append(mw, g.middleware...)
		mw = append(mw, m...)
	}
	g.b.Handle(endpoint, h, mw...)
}
