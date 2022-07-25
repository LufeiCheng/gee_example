package main

import (
	"fmt"
	"gee"
	"html/template"
	"net/http"
	"time"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	g := gee.Default()
	g.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	g.LoadHTMLGlob("templates/*")
	g.Static("/assets", "./static")

	g.GET("/", indexHandler)
	g.GET("/hello", helloHandler)
	g.POST("/login", loginHandler)
	g.GET("/hello/:name/go", helloNameHandler)
	g.GET("/panic", panicHandler)

	group1 := g.Group("/hello")
	group1.GET("/lufei", helloHandler)

	g.Run(":9999")
}

func indexHandler(c *gee.Context) {
	c.HTML(http.StatusOK, "css.html", nil)
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

func panicHandler(c *gee.Context) {
	names := []string{"geektutu"}
	c.String(http.StatusOK, names[100])
}
