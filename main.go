package main

import (
	"github.com/kataras/iris"
	"log"
)

func main() {
	// prefix log entries with date and time (with microseconds) e.g.
	// 2016/05/22 11:11:47.152342 location-api starting
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("location-api starting")

	api := iris.New()
	api.Get("/location", getLocation)
	api.Listen(":8080")
}

func getLocation(context *iris.Context) {
	context.Write("Lcation object returned here")
}
