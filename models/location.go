package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Location struct {
	Id                 bson.ObjectId `json:"id" bson:"_id"`
	Latitude           float64       `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude          float64       `json:"longitude,omitempty" bson:"longitude,omitempty"`
	Altitude           float64       `json:"alt,omitempty" bson:"alt,omitempty"`
	HorizontalAccuracy float64       `json:"horizontalAccuracy,omitempty" bson:"horizontalAccuracy,omitempty"`
	VerticalAccuracy   float64       `json:"verticalAccuracy,omitempty" bson:"verticalAccuracy,omitempty"`
	DeviceTime         time.Time     `json:"devicetime,omitempty" bson:"devicetime,omitempty"`
	Description        string        `json:"description,omitempty" bson:"description,omitempty"`
}

func (l *Location) IsValid() bool {
	if l.Latitude == 0 {
		return false
	}

	if l.Longitude == 0 {
		return false
	}

	return true
}

var TestLocation = Location{
	Id:                 bson.NewObjectId(),
	Latitude:           1.111,
	Longitude:          2.222,
	Altitude:           3.333,
	HorizontalAccuracy: 4.444,
	VerticalAccuracy:   5.555,
	DeviceTime:         time.Now(),
	Description:        "This is a valid Location object.",
}
