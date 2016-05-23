package main

import (
	"github.com/billglover/location-api/handlers"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"log"
)

func main() {
	var (
		dbName  = "test"
		address = ":8080"
	)

	// prefix log entries with date and time (with microseconds) e.g.
	// 2016/05/22 11:11:47.152342 location-api starting
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("location-api starting")

	// connect to the MongoDB database
	dbSession, err := mgo.Dial("192.168.99.100:32768")
	if err != nil {
		log.Panic(err)
	}
	defer dbSession.Close()

	// establish a new instance of the Iris server
	api := iris.New()

	// add a logging function to log requests
	api.UseFunc(handlers.LogHandler)

	// the following methods have been implemented
	api.Get("/location/:id", handlers.NewLocationHandler(dbSession, dbName).Get)
	api.Post("/location", handlers.NewLocationHandler(dbSession, dbName).Post)

	// the following methods are not implemented
	api.Get("/location", handlers.NotImplemented)
	api.Put("/location", handlers.NotImplemented)
	api.Delete("/location", handlers.NotImplemented)
	api.Head("/location", handlers.NotImplemented)
	api.Patch("/location", handlers.NotImplemented)
	api.Options("/location", handlers.NotImplemented)

	// start the server
	log.Printf("listening on %s", address)
	api.Listen(address)
}
