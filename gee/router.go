package gee

import (
	"log"
)

type router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {

	req := c.Req
	key := req.Method + "-" + req.URL.Path

	log.Printf("request: %s -> %s", req.Method, req.URL.Path)

	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(404, "404 NOT FOUND: %s \n", req.URL)
	}

}
