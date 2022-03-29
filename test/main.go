package main

import (
	"net/http"

	"github.com/hfpublic/myrouter"
)

func main() {
	r := myrouter.New()

	r.GET("/", func(c *myrouter.Context) {
		c.HTML(http.StatusOK, "<h1>hello myrouter</h1>")
	})

	r.GET("/hello", func(c *myrouter.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *myrouter.Context) {
		c.JSON(http.StatusOK, myrouter.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.RUN(":9090")
}
