#!/usr/bin/env bash

## Orders Service
docker build -t agarwalconsulting/rvstore-order-api:v1 --build-arg SERVICE=orders .
docker push agarwalconsulting/rvstore-order-api:v1

## Products Service
docker build -t agarwalconsulting/rvstore-product-api:v1 --build-arg SERVICE=products .
docker push agarwalconsulting/rvstore-product-api:v1

## Auth Service
docker build -t agarwalconsulting/rvstore-auth-api:v1 --build-arg SERVICE=auth .
docker push agarwalconsulting/rvstore-auth-api:v1

## Gateway Service
docker build -t agarwalconsulting/rvstore-api-gateway:v1 web/gateway-service
docker push agarwalconsulting/rvstore-api-gateway:v1

## Building UI
docker build -t agarwalconsulting/rvstore-ui:v1 web/rvstore
docker push agarwalconsulting/rvstore-ui:v1

## Building Order Simulator
docker build -t agarwalconsulting/rvstore-order-simulator:v1 tests/order-simulator
docker push agarwalconsulting/rvstore-order-simulator:v1
