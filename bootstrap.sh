#!/bin/sh

# create network
docker network create net

# start redis container
cd redis-host
docker-compose up --build &

cd ../customer-service
docker-compose up --build &

