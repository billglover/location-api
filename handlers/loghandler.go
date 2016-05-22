package handlers

import (
	"github.com/kataras/iris"
	"log"
)

func LogHandler(context *iris.Context) {
	// TODO: this log configuration has been duplicated from main.go
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	context.Next() // jumping ahead here allows us to log response information
	ip := context.RemoteAddr()
	status := context.Response.StatusCode()
	method := context.MethodString()
	path := context.PathString()

	// sample log format
	// 2016/05/22 13:43:10.968341 ::1 PATCH /location 501
	log.Println(ip, method, path, status)
}
