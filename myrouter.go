package myrouter

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRouter("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRouter("POST", pattern, handler)
}

func (engine *Engine) RUN(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	key := fmt.Sprintf("%s-%s", req.Method, req.URL.Path)
	if handler, has := engine.router[key]; has {
		handler(resp, req)
	} else {
		fmt.Fprintf(resp, "404 NOT FOUND: %s\n", req.URL.Path)
	}
}

func (engine *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	log.Printf("router %4s - %s", method, pattern)
	engine.router[key] = handler
}
