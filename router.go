package myrouter

import (
	"fmt"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlersChain
}

func newRouter() *router {
	return &router{roots: map[string]*node{}, handlers: make(map[string]HandlersChain)}
}

func (r *router) addRoute(method string, pattern string, handler HandlersChain) {
	if _, has := r.roots[method]; !has {
		r.roots[method] = &node{}
	}

	key := fmt.Sprintf("%s-%s", method, pattern)
	parts := parsePattern(pattern)
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := fmt.Sprintf("%s-%s", c.Method, n.pattern)
		c.handlers = r.handlers[key]
	} else {
		fmt.Fprintf(c.Writer, "404 NOT FOUND: %s\n", c.Path)
	}
	c.Next()
}

func (r *router) getRouter(method string, paths string) (*node, map[string]string) {
	if _, has := r.roots[method]; !has {
		return nil, nil
	}
	searchParts := parsePattern(paths)
	params := make(map[string]string)
	n := r.roots[method].search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for i := 0; i < len(parts); i++ {
			if parts[i][0] == ':' {
				params[parts[i][1:]] = searchParts[i]
			}
			if parts[i][0] == '*' && len(parts[i]) > 1 {
				params[parts[i][1:]] = strings.Join(searchParts[i:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func parsePattern(pattern string) []string {
	partArr := strings.Split(pattern, "/")
	parts := make([]string, 0)

	for _, part := range partArr {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}

	return parts
}
