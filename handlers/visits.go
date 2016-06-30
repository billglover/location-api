package handlers

import (
	"github.com/billglover/location-api/models"
	"github.com/gorilla/context"
	"gopkg.in/matryer/respond.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
	"log"
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
	respond.With(w, r, http.StatusNotImplemented, nil)
	context.Clear(r)
}

func VisitsPost(w http.ResponseWriter, r *http.Request) {
	respond.With(w, r, http.StatusNotImplemented, nil)
	context.Clear(r)
}
