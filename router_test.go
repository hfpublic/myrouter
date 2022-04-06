package myrouter

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/hello/:name"), []string{"hello", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/hello/*"), []string{"hello", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/hello/*name/*"), []string{"hello", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRouter(t *testing.T) {
	r := newTestRouter()
	node, params := r.getRouter("GET", "/hello/hfpublic")

	if node == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if node.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if params["name"] != "hfpublic" {
		t.Fatal("name should be equal to 'hfpublic'")
	}
	fmt.Printf("matched path: %s, params['name']: %s\n", node.pattern, params["name"])
}
