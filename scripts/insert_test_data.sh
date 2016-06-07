#!/bin/bash
mongo test --eval 'db.Locations.insert({_id: ObjectId("574cb30f4bf4c8f0c6a056e8"), latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "test location 1"});'
mongo test --eval 'db.Locations.insert({_id: ObjectId("574de23b5f810df11cad3498"), latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "test location 2"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "test location 3"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "test location 4"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "test location 5"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "test location 6"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "test location 7"});'
mongo test --eval 'db.Locations.insert({latitude: 1.1111, longitude: 2.2222, altitude: 3.3333, horizontalAccuracy: 4.4444, verticalAccuracy: 5.5555, description: "test location 8"});'