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

	r.RUN(":9090")
}
