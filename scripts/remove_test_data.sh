#!/bin/bash
mongo test --eval 'db.Locations.drop();'
mongo test --eval 'db.Visits.drop();'