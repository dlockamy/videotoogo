#!/bin/bash
docker stop $(docker ps -a | grep "www" | awk '{print $1}')
docker stop $(docker ps -a | grep "upload" | awk '{print $1}')
docker stop $(docker ps -a | grep "listen" | awk '{print $1}')
docker stop $(docker ps -a | grep "processor" | awk '{print $1}')
