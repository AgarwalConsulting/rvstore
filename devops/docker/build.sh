#!/usr/bin/env bash

## Gateway Service
docker build -t agarwalconsulting/rvstore-api-gateway:v1 web/gateway-service
docker push agarwalconsulting/rvstore-api-gateway:v1

## Building UI
docker build -t agarwalconsulting/rvstore-ui:v1 web/rvstore
docker push agarwalconsulting/rvstore-ui:v1

## Building Order Simulator
docker build -t agarwalconsulting/rvstore-order-simulator:v1 tests/order-simulator
docker push agarwalconsulting/rvstore-order-simulator:v1
