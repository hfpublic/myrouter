package myrouter

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

type RouterGroup struct {
	prefix      string
	middlewares HandlersChain
	engine      *Engine
	parent      *RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: fmt.Sprintf("%s%s", group.prefix, prefix),
		engine: engine,
		parent: group,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (engine *Engine) RUN(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContest(w, r)
	engine.router.handle(c)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := fmt.Sprintf("%s%s", group.prefix, comp)
	handlers := make(HandlersChain, 0)
	groupPath := group
	for groupPath != nil {
		handlers = append(groupPath.middlewares, handlers...)
		groupPath = groupPath.parent
	}
	handlers = append(handlers, handler)
	log.Printf("router %4s-%s", method, pattern)
	group.engine.router.addRoute(method, pattern, handlers)
}
