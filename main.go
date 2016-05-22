package main

import (
	"github.com/billglover/location-api/handlers"
	"github.com/kataras/iris"
	"log"
)

func main() {
	// prefix log entries with date and time (with microseconds) e.g.
	// 2016/05/22 11:11:47.152342 location-api starting
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("location-api starting")

	api := iris.New()

	// log all incomming requests
	api.UseFunc(func(context *iris.Context) {
		log.Println(context.MethodString(), context.PathString())
		context.Next()
	})

	// the following methods have been implemented
	api.Get("/location", handlers.GetLocation)

	// the following methods are not implemented
	api.Get("/location", handlers.NotImplemented)
	api.Post("/location", handlers.NotImplemented)
	api.Put("/location", handlers.NotImplemented)
	api.Delete("/location", handlers.NotImplemented)
	api.Head("/location", handlers.NotImplemented)
	api.Patch("/location", handlers.NotImplemented)
	api.Options("/location", handlers.NotImplemented)
	api.Listen(":8080")
}
