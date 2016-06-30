#!/bin/bash
./scripts/remove_test_data.sh
./scripts/insert_test_data.sh
go test -v ./...
./scripts/remove_test_data.sh
