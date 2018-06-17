#!/bin/bash

cd cmd/listen
env GOOS=linux  go build -o listen main.go
docker build -t local/listen .
rm listen
cd ../..

cd cmd/upload
env GOOS=linux  go build -o upload main.go 
docker build -t local/upload .
rm upload
cd ../..

cd cmd/processor
env GOOS=linux  go build -o processor main.go 
docker build -t local/processor .
rm processor
cd ../..

cd cmd/web
env GOOS=linux  go build -o web main.go 
cp -fr ../../www ./www
docker build -t local/www .
rm web
rm -fr ./www 
cd ../..
