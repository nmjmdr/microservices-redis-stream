#!/bin/sh

# create network
docker network create net

# start redis container
cd redis-host
docker-compose up --build &


# start db-host
cd ..
cd db-host
docker-compose up --build &

# start account-service
cd ..
cd account-service
docker-compose up --build &

# start customer-service
cd ..
cd customer-service
docker-compose up --build &

