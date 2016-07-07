package handlers

import (
	"github.com/billglover/location-api/models"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/matryer/respond.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
	"log"
	"io/ioutil"
	"encoding/json"
)

func VisitsGet(w http.ResponseWriter, r *http.Request) {
	page, _     := strconv.Atoi(r.URL.Query().Get("page"))		// on error page is set to 0
	per_page, _ := strconv.Atoi(r.URL.Query().Get("per_page"))	// on error per_page is set to 0
	time_from_string := r.URL.Query().Get("time_from")
	time_to_string := r.URL.Query().Get("time_to")

	// construct a search query
	f := bson.M{}
	hAccuracy := bson.M{"$lt": 100}
	f["horizontalAccuracy"] = hAccuracy
	f["description"] = "visit"

	if time_from_string != "" {
		time_from, _ := time.Parse(time.RFC3339, time_from_string)
		time_range := bson.M{}
		time_range["$gte"] = time_from
		f["arrivalTime"] = time_range
	}
	if time_to_string != "" {
		time_to, _ := time.Parse(time.RFC3339, time_to_string)
		time_range := bson.M{}
		time_range["$lte"] = time_to
		f["departureTime"] = time_range
	}

	db := context.Get(r, "db").(*mgo.Session)
	l := []models.Visit{}
	err := db.DB("").C("Visits").Find(f).Sort("-_id").Skip((page-1)*per_page).Limit(per_page).All(&l)

	if err != nil {
		log.Println(err.Error())
		respond.WithStatus(w, r, http.StatusInternalServerError)
		return
	}
	respond.With(w, r, http.StatusOK, l)
	context.Clear(r)
}

func VisitsGetOne(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*mgo.Session)

	vars := mux.Vars(r)
	id := vars["id"]
	if !bson.IsObjectIdHex(id) {
		respond.WithStatus(w, r, http.StatusNotFound)
		return
	}

	v := &models.Visit{}
	err := db.DB("").C("Visits").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&v)
	if err != nil {
		log.Println("VisitsGetOne", id, err.Error())
		if err.Error() == "not found" {
			respond.WithStatus(w, r, http.StatusNotFound)
			return
		} else {
			respond.WithStatus(w, r, http.StatusInternalServerError)
			return	
		}		
	}
	respond.With(w, r, http.StatusOK, v)
	context.Clear(r)
}

func VisitsPost(w http.ResponseWriter, r *http.Request) {
	var visits []models.Location
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		log.Println(errBody)
		respond.WithStatus(w, r, http.StatusInternalServerError)
		return
	}

	errJson := json.Unmarshal(body, &visits)
	if errJson != nil {
		//log.Printf("Unable to convert body to valid Location object. Received: %s", visits)
		respond.WithStatus(w, r, http.StatusBadRequest)
		return
	}

	db := context.Get(r, "db").(*mgo.Session)

	for i, _ := range visits {
		if visits[i].IsInvalid() {
			//log.Printf("Unable to convert body to valid Location object. Received: %s", visits)
			respond.WithStatus(w, r, http.StatusBadRequest)
			return		
		}
		visits[i].Id = bson.NewObjectId()
		errDb := db.DB("").C("Visits").Insert(visits[i])
		if errDb != nil {
			log.Println(errDb)
			respond.WithStatus(w, r, http.StatusInternalServerError)
			return
		}
	}

	respond.With(w, r, http.StatusCreated, visits)
	context.Clear(r)
}
