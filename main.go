package main

import (
	"github.com/billglover/location-api/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		dbName  = "test"
		address = ":8080"
		dbUrl   = "localhost:27017"
	)

	if os.Getenv("DB_URL") != "" {
		dbUrl = os.Getenv("DB_URL")
	}

	// prefix log entries with date and time (with microseconds) e.g.
	// 2016/05/22 11:11:47.152342 location-api starting
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("location-api starting")
	log.Println("connecting to MongoDB:", dbUrl)
	log.Println("using DB:", dbName)

	// connect to the MongoDB database
	dbSession, err := mgo.Dial(dbUrl)
	if err != nil {
		log.Panic(err)
	}
	defer dbSession.Close()

	router := apiRouter(dbSession)

	log.Printf("listening on %s", address)
	log.Fatalln(http.ListenAndServe(address, router))

}

func apiRouter(d *mgo.Session) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/locations/{id}", db.WithDB(d, handlers.LocationsGet)).Methods("GET")
	router.HandleFunc("/locations", db.WithDB(d, handlers.LocationsPost)).Methods("POST")
	return router
}

