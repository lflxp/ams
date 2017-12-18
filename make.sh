#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
docker stop server 
docker rm server 
docker build -t lxp/ams . && docker run -d --name server -p 8080:8080 -p 8088:8088 lxp/ams

