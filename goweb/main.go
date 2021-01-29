package main

/*
(1) index
curl -i http://localhost:9999/index
HTTP/1.1 200 OK
Date: Sun, 01 Sep 2019 08:12:23 GMT
Content-Length: 19
Content-Type: text/html; charset=utf-8
<h1>Index Page</h1>

(2) v1
$ curl -i http://localhost:9999/v1/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 18:11:07 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello Gee</h1>

(3)
$ curl "http://localhost:9999/v1/hello?name=geektutu"
hello geektutu, you're at /v1/hello

(4)
$ curl "http://localhost:9999/v2/hello/geektutu"
hello geektutu, you're at /hello/geektutu

(5)
$ curl "http://localhost:9999/v2/login" -X POST -d 'username=geektutu&password=1234'
{"password":"1234","username":"geektutu"}

(6)
$ curl "http://localhost:9999/hello"
404 NOT FOUND: /hello
*/

import (
	"fmt"
	"gee"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsData(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	//r := gee.New()
	//r.Use(gee.Logger()) // global middleware
	//r.SetFuncMap(template.FuncMap{
	//	"FormatAsDate": FormatAsData,
	//})
	//r.LoadHTMLGlob("./templates/*")
	//r.Static("/assets", "./static")
	//
	//stu1 := &student{Name: "Geektutu", Age: 20}
	//stu2 := &student{Name: "Jack", Age: 22}
	//r.GET("/", func(c *gee.Context) {
	//	c.HTML(http.StatusOK, "css.tmpl", nil)
	//})
	//r.GET("/students", func(c *gee.Context) {
	//	c.HTML(http.StatusOK, "arr.tmpl", gee.H{
	//		"title":  "gee",
	//		"stuArr": [2]*student{stu1, stu2},
	//	})
	//})
	//r.GET("/date", func(c *gee.Context) {
	//	c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
	//		"title": "gee",
	//		"now":   time.Date(2021, 1, 28, 0, 0, 0, 0, time.UTC),
	//	})
	//})

	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	r.GET("panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":9999")
}
