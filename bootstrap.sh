#!/bin/sh

# create network
docker network create net

# start db-host
cd db-host
docker-compose up --build &

# Ideally should poll the container to check if its up
sleep 10

# start redis container
cd ..
cd redis-host
docker-compose up --build &

# Ideally should poll the container to check if its up
sleep 10

# start customer-service
cd ..
cd customer-service
docker-compose up --build &

# start account-service
cd ..
cd account-service
docker-compose up --build &



