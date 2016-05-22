package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Location struct {
	Id                 bson.ObjectId `json:"id" bson:"_id"`
	Latitude           float32       `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude          float32       `json:"longitude,omitempty" bson:"longitude,omitempty"`
	Altitude           float32       `json:"alt,omitempty" bson:"alt,omitempty"`
	HorizontalAccuracy float32       `json:"horizontalAccuracy,omitempty" bson:"horizontalAccuracy,omitempty"`
	VerticalAccuracy   float32       `json:"verticalAccuracy,omitempty" bson:"verticalAccuracy,omitempty"`
	DeviceTime         time.Time     `json:"devicetime,omitempty" bson:"devicetime,omitempty"`
	Description        string        `json:"description,omitempty" bson:"description,omitempty"`
}
