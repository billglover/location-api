package handlers

import (
	"encoding/json"
	"github.com/billglover/location-api/models"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/matryer/respond.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func LocationsGet(w http.ResponseWriter, r *http.Request) {
	page, _     := strconv.Atoi(r.URL.Query().Get("page"))		// on error page is set to 0
	per_page, _ := strconv.Atoi(r.URL.Query().Get("per_page"))	// on error per_page is set to 0

	db := context.Get(r, "db").(*mgo.Session)
	l := []models.Location{}
	err := db.DB("test").C("Locations").Find(bson.M{}).Skip(page*per_page).Limit(per_page).All(&l)
	if err != nil {
		log.Println(err.Error())
		respond.WithStatus(w, r, http.StatusInternalServerError)
		return
	}
	respond.With(w, r, http.StatusOK, l)
	context.Clear(r)
}

func LocationsGetOne(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*mgo.Session)

	vars := mux.Vars(r)
	id := vars["id"]
	if !bson.IsObjectIdHex(id) {
		respond.WithStatus(w, r, http.StatusNotFound)
		return
	}

	l := &models.Location{}
	err := db.DB("test").C("Locations").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&l)
	if err != nil {
		log.Println(err.Error())
		respond.WithStatus(w, r, http.StatusInternalServerError)
		return
	}
	respond.With(w, r, http.StatusOK, l)
	context.Clear(r)
}

func LocationsPost(w http.ResponseWriter, r *http.Request) {
	l := models.Location{}
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		log.Println(errBody)
		respond.WithStatus(w, r, http.StatusInternalServerError)
		return
	}

	errJson := json.Unmarshal(body, &l)
	if errJson != nil || l.IsInvalid() {
		log.Printf("Unable to convert body to valid Location object. Received: %s", l)
		respond.WithStatus(w, r, http.StatusBadRequest)
		return
	}

	// give our new Location an ID
	l.Id = bson.NewObjectId()

	db := context.Get(r, "db").(*mgo.Session)
	errDb := db.DB("test").C("Locations").Insert(l)
	if errDb != nil {
		log.Println(errDb)
		respond.WithStatus(w, r, http.StatusInternalServerError)
		return
	}

	respond.With(w, r, http.StatusCreated, l)
}
