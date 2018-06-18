#!/bin/bash
echo "Starting services. Use killservices.sh to exit all services."

# Check to make sure we have an active database
DATAROOT=./var
BLOCKSDIR=./var/blocks
DATABASEDIR=./var/data
DATABASE=./var/data/videos.json
UPLOADDIR=./var/uploads


if [ ! -d $DATAROOT ]; then
   mkdir $DATAROOT;
fi

if [ ! -d $BLOCKSDIR ]; then
   mkdir $BLOCKSDIR;
fi

if [ ! -d $DATABASEDIR ]; then
   mkdir $DATABASEDIR;
fi
   
if [ ! -f $DATABASE ]; then
   go run cmd/builddb/buildDB.go
fi

if [ ! -d $UPLOADDIR ]; then
   mkdir $UPLOADDIR;
fi

ROOT_DIR=$(pwd)

echo "root = $ROOT_DIR"
docker run -d -p 3000:3000 local/www
docker run -d -p 3001:3001 -v $ROOT_DIR/var:/var local/upload
docker run -d -p 3002:3002 -v $ROOT_DIR/var:/var local/listen
docker run -d -v $ROOT_DIR/var:/var local/processor
