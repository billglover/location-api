package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Visit struct {
	Id                 bson.ObjectId `json:"id" bson:"_id"`
	Latitude           float64       `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude          float64       `json:"longitude,omitempty" bson:"longitude,omitempty"`
	HorizontalAccuracy float64       `json:"horizontalAccuracy,omitempty" bson:"horizontalAccuracy,omitempty"`
	ArrivalTime        time.Time     `json:"arrivalTime,omitempty" bson:"arrivalTime,omitempty"`
	DepartureTime      time.Time     `json:"departureTime,omitempty" bson:"departureTime,omitempty"`
	Description        string        `json:"description,omitempty" bson:"description,omitempty"`
}

func (l *Visit) IsValid() bool {
	if l.Latitude           == 0 { return false }
	if l.Longitude          == 0 { return false }
	if l.HorizontalAccuracy == 0 { return false }
	if l.ArrivalTime.IsZero()    { return false }
	if l.DepartureTime.IsZero()  { return false }
	if l.Description != "visit" { return false }

	return true
}

func (l *Visit) IsInvalid() bool {
	return !l.IsValid()
}
