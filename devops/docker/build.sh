#!/usr/bin/env bash

## Building UI
docker build -t agarwalconsulting/rvstore-ui:v1 web/rvstore
docker push agarwalconsulting/rvstore-ui:v1

## Building Order Simulator
docker build -t agarwalconsulting/rvstore-order-simulator:v1 tests/order-simulator
docker push agarwalconsulting/rvstore-order-simulator:v1
