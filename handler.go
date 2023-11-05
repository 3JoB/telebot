package crare

// Handle stores each handler and its corresponding middleware,
// which may be optimized in the future.
type Handle struct {
	Do         HandlerFunc
	Middleware []HandlerFunc
}

// HandlerFunc represents a handler function, which is
// used to handle actual endpoints.
type HandlerFunc func(*Context) error

// Execute handler
func (h *Handle) do(c *Context) error {
	if len(h.Middleware) > 0 {
		if !c.next {
			return nil
		}
	}
	return h.Do(c)
}

// Execution middleware
func (h *Handle) doMiddleware(c *Context) error {
	if len(h.Middleware) > 0 {
		for i, r := range h.Middleware {
			if !c.next && i != 0 {
				return nil
			}
			c.next = false
			if err := r(c); err != nil {
				return err
			}
		}
	}
	return nil
}
