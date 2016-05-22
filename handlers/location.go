package handlers

import (
	"github.com/billglover/location-api/models"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// The LocationHandler struct allows us to pass in database session details
type LocationHandler struct {
	session        *mgo.Session
	dbName         string
	collectionName string
}

func NewLocationHandler(s *mgo.Session, d string) *LocationHandler {
	return &LocationHandler{session: s, dbName: d, collectionName: "Locations"}
}

func (lh LocationHandler) Get(context *iris.Context) {

	// check if we have a valid ID, if not return 404
	if !bson.IsObjectIdHex(context.Param("id")) {
		context.SetStatusCode(iris.StatusNotFound)
		context.Write("Invalid Location ID")
		return
	}
	id := bson.ObjectIdHex(context.Param("id"))

	// create an empty Location
	location := models.Location{}

	// create a copy of the DB session and close once we are complete
	s := lh.session.Copy()
	c := s.DB(lh.dbName).C(lh.collectionName)
	defer s.Close()

	// search for our Location record by ID
	err := c.FindId(id).One(&location)
	if err != nil {

		// if we don't find anything return 404
		// otherwise something else went wrong
		if err.Error() == "not found" {
			context.SetStatusCode(iris.StatusNotFound)
			context.Write(iris.StatusText(iris.StatusNotFound))
		} else {
			context.SetStatusCode(iris.StatusInternalServerError)
			context.Write(iris.StatusText(iris.StatusInternalServerError))
			log.Fatal(err)
		}
		return
	}

	// if we found a Location object return it as JSON
	context.JSON(iris.StatusOK, location)
}
