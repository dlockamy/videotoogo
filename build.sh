#!/bin/bash

cd listen
env GOOS=linux  go build -o listen main.go
docker build -t local/listen .
rm listen
cd ..

cd upload
env GOOS=linux  go build -o upload main.go 
docker build -t local/upload .
rm upload
cd ..

cd processor
env GOOS=linux  go build -o upload main.go 
docker build -t local/processor .
rm processor
cd ..

cd webclient
env GOOS=linux  go build -o webclient main.go 
cp -fr ../www ./
docker build -t local/webclient .
rm webclient
rm -fr ./www 
cd ..