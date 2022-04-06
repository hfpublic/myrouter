package main

import (
	"net/http"

	"github.com/hfpublic/myrouter"
)

func main() {
	r := myrouter.New()

	// curl -i http://localhost:9090/
	r.GET("/", func(c *myrouter.Context) {
		c.HTML(http.StatusOK, "<h1>hello myrouter</h1>")
	})

	// curl "http://localhost:9090/hello?name=hfpublic"
	r.GET("/hello", func(c *myrouter.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
	})

	// curl "http://localhost:9090/login" -X POST -d 'username=hfpublic&password=1234'
	r.POST("/login", func(c *myrouter.Context) {
		c.JSON(http.StatusOK, myrouter.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	// curl "http://localhost:9090/hello/hfpublic"
	r.GET("/hello/:name", func(c *myrouter.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Param("name"), c.Path)
	})

	// curl "http://localhost:9090/assets/css/hfpublic.css"
	r.GET("/assets/*filepath", func(c *myrouter.Context) {
		c.JSON(http.StatusOK, myrouter.H{"filepath": c.Param("filepath")})
	})

	v1 := r.Group("v1")
	{
		// curl "http://localhost:9090/v1/index"
		v1.GET("/index", func(ctx *myrouter.Context) {
			ctx.HTML(http.StatusOK, "<h1>hello myrouter, v1</h1>")
		})
	}

	v2 := r.Group("v2")
	{
		// curl "http://localhost:9090/v2/index"
		v2.GET("/index", func(ctx *myrouter.Context) {
			ctx.HTML(http.StatusOK, "<h1>hello myrouter, v2</h1>")
		})
		// curl "http://localhost:9090/v2/hello/hfpublic"
		v2.GET("/hello/:name", func(c *myrouter.Context) {
			c.String(http.StatusOK, "v2 hello %s, you are at %s\n", c.Param("name"), c.Path)
		})
	}

	r.RUN(":9090")
}
