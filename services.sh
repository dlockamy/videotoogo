#!/bin/bash
echo "Starting services. Use killservices.sh to exit all services."
go run webclient/main.go &
go run listen/main.go &
go run upload/main.go &
