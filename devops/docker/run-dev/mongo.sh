#!/usr/bin/env bash

docker run --rm -it -p 27017:27017 \
  -e MONGO_INITDB_ROOT_USERNAME=root \
  -e MONGO_INITDB_ROOT_PASSWORD=example \
  mongo

# docker run --rm -it -p 8081:8081 \
#   -e ME_CONFIG_MONGODB_ADMINUSERNAME=root \
#   -e ME_CONFIG_MONGODB_ADMINPASSWORD=example \
#   -e ME_CONFIG_MONGODB_SERVER=localhost \
#   mongo-express
