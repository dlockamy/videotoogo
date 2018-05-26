#!/bin/bash
echo "Starting services. Use killservices.sh to exit all services."

ROOT_DIR=$(pwd)

echo "root = $ROOT_DIR"
docker run -d -p 3000:3000 local/webclient
docker run -d -p 3001:3001 -v $ROOT_DIR/data:/data -v $ROOT_DIR/blocks:/blocks local/upload
docker run -d -p 3002:3002 -v $ROOT_DIR/data:/data -v $ROOT_DIR/blocks:/blocks local/listen