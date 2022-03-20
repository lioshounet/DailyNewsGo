#!/bin/bash

app_deploy(){
	echo "deploy ..."
  docker stop thor_web
  docker rm thor_web
  docker rmi thor_go:v1.0
	docker-compose up -d
	echo "done"
}

tree ./
app_deploy