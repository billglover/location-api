#!/bin/bash
mongo test --eval 'db.Locations.insert({_id: ObjectId("574cb30f4bf4c8f0c6a056e8"), latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "location"});'
mongo test --eval 'db.Locations.insert({_id: ObjectId("574de23b5f810df11cad3498"), latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "location"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "location"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "location"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "location"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "location"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "location"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "location"});'


mongo test --eval 'db.Visits.insert({latitude: 1.1111, longitude: 2.2222, horizontalAccuracy: 3.3333, arrivalTime: new ISODate("2016-06-01T06:00:00Z"), departureTime: new ISODate("2016-06-01T07:00:00Z"), description: "visit"});'
mongo test --eval 'db.Visits.insert({latitude: 1.1111, longitude: 2.2222, horizontalAccuracy: 3.3333, arrivalTime: new ISODate("2016-06-02T06:00:00Z"), departureTime: new ISODate("2016-06-02T07:00:00Z"), description: "visit"});'
mongo test --eval 'db.Visits.insert({latitude: 1.1111, longitude: 2.2222, horizontalAccuracy: 3.3333, arrivalTime: new ISODate("2016-06-03T06:00:00Z"), departureTime: new ISODate("2016-06-03T07:00:00Z"), description: "visit"});'