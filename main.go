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

	// add a logging function to log requests
	api.UseFunc(handlers.LogHandler)

	// the following methods have been implemented
	api.Get("/location", handlers.GetLocation)

	// the following methods are not implemented
	api.Post("/location", handlers.NotImplemented)
	api.Put("/location", handlers.NotImplemented)
	api.Delete("/location", handlers.NotImplemented)
	api.Head("/location", handlers.NotImplemented)
	api.Patch("/location", handlers.NotImplemented)
	api.Options("/location", handlers.NotImplemented)

	// start the server
	log.Println("listening on :8080")
	api.Listen(":8080")
}
