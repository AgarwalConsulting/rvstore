#!/usr/bin/env bash

docker kill mongodb
docker rm mongodb

docker kill rvstore-order-api
docker rm rvstore-order-api

docker kill rvstore-product-api
docker rm rvstore-product-api

docker kill rvstore-auth-api
docker rm rvstore-auth-api

docker kill rvstore-api-gateway
docker rm rvstore-api-gateway

docker kill rvstore-ui
docker rm rvstore-ui

docker kill rvstore-order-simulator
docker rm rvstore-order-simulator
