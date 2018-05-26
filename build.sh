#!/bin/bash

cd listen
go build -o listen main.go
docker build -t local/listen .
rm listen
cd ..

cd upload
go build -o upload main.go 
docker build -t local/upload .
rm upload
cd ..

cd webclient
go build -o webclient main.go 
cp -fr ../www ./
docker build -t webclient .
rm webclient
rm -fr ./www 
cd ..