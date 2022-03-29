package main

import (
	"fmt"
	"net/http"

	"github.com/hfpublic/myrouter"
)

func main() {
	r := myrouter.New()

	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	})

	r.RUN(":9090")
}
