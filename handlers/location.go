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
	"time"
)

func today() time.Time {
	t := time.Now()
    year, month, day := t.Date()
    return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func LocationsGet(w http.ResponseWriter, r *http.Request) {
	page, _     := strconv.Atoi(r.URL.Query().Get("page"))		// on error page is set to 0
	per_page, _ := strconv.Atoi(r.URL.Query().Get("per_page"))	// on error per_page is set to 0

	db := context.Get(r, "db").(*mgo.Session)
	l := []models.Location{}
	err := db.DB("").C("Locations").Find(bson.M{"horizontalAccuracy": bson.M{"$lt": 100}, "description": "location", "devicetime": bson.M{"$gte": today()}}).Sort("-_id").Skip((page-1)*per_page).Limit(per_page).All(&l)

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
	err := db.DB("").C("Locations").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&l)
	if err != nil {
		log.Println("LocationGetOne", id, err.Error())
		if err.Error() == "not found" {
			respond.WithStatus(w, r, http.StatusNotFound)
			return
		} else {
			respond.WithStatus(w, r, http.StatusInternalServerError)
			return	
		}		
	}
	respond.With(w, r, http.StatusOK, l)
	context.Clear(r)
}

func LocationsPost(w http.ResponseWriter, r *http.Request) {
	var ls []models.Location
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		log.Println(errBody)
		respond.WithStatus(w, r, http.StatusInternalServerError)
		return
	}

	errJson := json.Unmarshal(body, &ls)
	if errJson != nil {
		//log.Printf("Unable to convert body to valid Location object. Received: %s", ls)
		respond.WithStatus(w, r, http.StatusBadRequest)
		return
	}

	db := context.Get(r, "db").(*mgo.Session)

	for i, _ := range ls {
		if ls[i].IsInvalid() {
			//log.Printf("Unable to convert body to valid Location object. Received: %s", ls)
			respond.WithStatus(w, r, http.StatusBadRequest)
			return		
		}
		ls[i].Id = bson.NewObjectId()
		errDb := db.DB("").C("Locations").Insert(ls[i])
		if errDb != nil {
			log.Println(errDb)
			respond.WithStatus(w, r, http.StatusInternalServerError)
			return
		}
	}

	respond.With(w, r, http.StatusCreated, ls)
}
