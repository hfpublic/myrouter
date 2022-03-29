package myrouter

import (
	"fmt"
	"log"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	log.Printf("router %4s - %s", method, pattern)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := fmt.Sprintf("%s-%s", c.Method, c.Path)
	if handler, has := r.handlers[key]; has {
		handler(c)
	} else {
		fmt.Fprintf(c.Writer, "404 NOT FOUND: %s\n", c.Path)
	}
}
