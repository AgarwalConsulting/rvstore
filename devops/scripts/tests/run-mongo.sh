#!/usr/bin/env bash

docker run --rm -itd -p 27017:27017 \
  --name mongodb \
  -e MONGO_INITDB_ROOT_USERNAME=root \
  -e MONGO_INITDB_ROOT_PASSWORD=example \
  mongo
