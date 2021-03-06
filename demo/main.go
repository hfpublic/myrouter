package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/hfpublic/myrouter"
)

func main() {
	r := myrouter.New()

	// curl -i http://localhost:9090/
	// r.GET("/", func(c *myrouter.Context) {
	// 	c.HTML(http.StatusOK, "<h1>hello myrouter</h1>")
	// })

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
			ctx.JSON(http.StatusOK, myrouter.H{
				"message": "hello myrouter, v1",
			})
		})
	}

	v2 := r.Group("v2")
	{
		// curl "http://localhost:9090/v2/index"
		v2.GET("/index", func(ctx *myrouter.Context) {
			ctx.JSON(http.StatusOK, myrouter.H{
				"message": "hello myrouter, v2",
			})
		})
		v2.Use(myrouter.Logger())
		// curl "http://localhost:9090/v2/hello/hfpublic"
		v2.GET("/hello/:name", func(c *myrouter.Context) {
			c.String(http.StatusOK, "v2 hello %s, you are at %s\n", c.Param("name"), c.Path)
		})
	}

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	// curl "http://localhost:9090/assets/testfile.txt"
	r.Static("/assets", "./static")

	// "http://localhost:9090/tmpl/css"
	r.GET("/tmpl/css", func(ctx *myrouter.Context) {
		ctx.HTML(http.StatusOK, "css.tmpl", nil)
		ctx.Next()
	})

	aaa := &testStruct{Name: "aaa", Age: 20}
	bbb := &testStruct{Name: "bbb", Age: 22}
	// "http://localhost:9090/tmpl/struct"
	r.GET("/tmpl/struct", func(ctx *myrouter.Context) {
		ctx.HTML(http.StatusOK, "arr.tmpl", myrouter.H{
			"title":   "myrouter",
			"structs": []*testStruct{aaa, bbb},
		})
	})

	// "http://localhost:9090/tmpl/func"
	r.GET("/tmpl/func", func(ctx *myrouter.Context) {
		ctx.HTML(http.StatusOK, "func.tmpl", myrouter.H{
			"title": "myrouter",
			"now":   time.Now().UTC(),
		})
	})

	r.Use(myrouter.Recovery())

	// curl "http://localhost:9090/panic"
	r.GET("/panic", func(ctx *myrouter.Context) {
		names := []string{"hfpublic"}
		ctx.String(http.StatusOK, names[100])
	})

	r.RUN(":9090")
}

type testStruct struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
