#!/bin/bash
docker rm $(docker ps -a | grep "web" | awk '{print $1}')
docker rm $(docker ps -a | grep "upload" | awk '{print $1}')
docker rm $(docker ps -a | grep "listen" | awk '{print $1}')
docker rm $(docker ps -a | grep "processor" | awk '{print $1}')
