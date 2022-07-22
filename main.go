package main

import (
	"gee"
	"net/http"
)

func main() {
	g := gee.New()
	g.GET("/", indexHandler)
	g.GET("/hello", helloHandler)
	g.POST("/login", loginHandler)
	g.GET("/hello/:name/go", helloNameHandler)
	g.Run(":9999")
}

func indexHandler(c *gee.Context) {
	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
}

func helloHandler(c *gee.Context) {
	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
}

func loginHandler(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}

func helloNameHandler(c *gee.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "<h1>Hello %s go</h1>", name)
}
