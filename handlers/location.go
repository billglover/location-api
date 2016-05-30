package handlers

import (
	"github.com/billglover/location-api/models"
	"github.com/gorilla/context"
	"gopkg.in/matryer/respond.v1"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

func LocationsGet(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*mgo.Session)
	l := &models.Location{}
	err := db.DB("test").C("Locations").Find(nil).One(&l)
	if err != nil {
		log.Println(err.Error())
		respond.WithStatus(w, r, http.StatusInternalServerError)
		return
	}
	respond.With(w, r, http.StatusOK, l)
	context.Clear(r)
}

func LocationsPost(w http.ResponseWriter, r *http.Request) {
	respond.WithStatus(w, r, http.StatusNotImplemented)
}

// The LocationHandler struct allows us to pass in database session details
// type LocationHandler struct {
// 	session        *mgo.Session
// 	dbName         string
// 	collectionName string
// }

// func NewLocationHandler(s *mgo.Session, d string) *LocationHandler {
// 	return &LocationHandler{session: s, dbName: d, collectionName: "Locations"}
// }

// func (lh LocationHandler) Get(context *iris.Context) {

// 	// check if we have a valid ID, if not return 404
// 	if !bson.IsObjectIdHex(context.Param("id")) {
// 		context.SetStatusCode(iris.StatusNotFound)
// 		context.Write("Invalid Location ID")
// 		return
// 	}
// 	id := bson.ObjectIdHex(context.Param("id"))

// 	// create an empty Location
// 	location := models.Location{}

// 	// create a copy of the DB session and close once we are complete
// 	s := lh.session.Copy()
// 	c := s.DB(lh.dbName).C(lh.collectionName)
// 	defer s.Close()

// 	// search for our Location record by ID
// 	err := c.FindId(id).One(&location)
// 	if err != nil {

// 		// if we don't find anything return 404
// 		// otherwise something else went wrong
// 		if err.Error() == "not found" {
// 			context.SetStatusCode(iris.StatusNotFound)
// 			context.Write(iris.StatusText(iris.StatusNotFound))
// 		} else {
// 			context.SetStatusCode(iris.StatusInternalServerError)
// 			context.Write(iris.StatusText(iris.StatusInternalServerError))
// 			log.Fatal(err)
// 		}
// 		return
// 	}

// 	// if we found a Location object return it as JSON
// 	context.JSON(iris.StatusOK, location)
// }

// func (lh LocationHandler) Post(context *iris.Context) {
// 	// get a location from the request body
// 	location := models.Location{}
// 	err := context.ReadJSON(&location)
// 	if err != nil {
// 		context.JSON(iris.StatusBadRequest, iris.StatusText(iris.StatusBadRequest))
// 		return
// 	}

// 	// create an ID
// 	location.Id = bson.NewObjectId()

// 	// confirm it is valid
// 	if !location.IsValid() {
// 		context.JSON(iris.StatusBadRequest, iris.StatusText(iris.StatusBadRequest))
// 		return
// 	}

// 	// submit to the DB
// 	s := lh.session.Copy()
// 	c := s.DB(lh.dbName).C(lh.collectionName)
// 	defer s.Close()
// 	err = c.Insert(&location)
// 	if err != nil {
// 		log.Println(err)
// 		context.JSON(iris.StatusInternalServerError, iris.StatusText(iris.StatusInternalServerError))
// 		return
// 	}

// 	// return the ID if successful
// 	log.Println(location)
// 	context.JSON(iris.StatusOK, location)
// }
