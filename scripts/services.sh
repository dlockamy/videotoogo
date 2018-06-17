#!/bin/bash
echo "Starting services. Use killservices.sh to exit all services."

ROOT_DIR=$(pwd)

echo "root = $ROOT_DIR"
docker run -d -p 3000:3000 local/www
docker run -d -p 3001:3001 -v $ROOT_DIR/var:/var local/upload
docker run -d -p 3002:3002 -v $ROOT_DIR/var:/var local/listen
docker run -d -v $ROOT_DIR/var:/var local/processor
