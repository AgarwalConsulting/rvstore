#!/usr/bin/env bash

docker volume create mongodata
docker network create rvstore

docker run --rm -itd -p 27017:27017 \
  --name mongodb \
  --network rvstore \
  --network-alias mongodb \
  -e MONGO_INITDB_ROOT_USERNAME=root \
  -e MONGO_INITDB_ROOT_PASSWORD=example \
  -v mongodata:/data/db \
  mongo

docker run --rm -d -p 9002:9002 \
  --name rvstore-order-api \
  --network rvstore \
  --network-alias rvstore-order-api \
  -e MONGO_DB_URL="mongodb://root:example@mongodb/admin" \
  agarwalconsulting/rvstore-order-api:v1

docker run --rm -d -p 9001:9001 \
  --name rvstore-product-api \
  --network rvstore \
  --network-alias rvstore-product-api \
  agarwalconsulting/rvstore-product-api:v1

docker run --rm -d -p 9003:9003 \
  --name rvstore-auth-api \
  --network rvstore \
  --network-alias rvstore-auth-api \
  agarwalconsulting/rvstore-auth-api:v1

docker run --rm -d -p 9000:9000 \
  --name rvstore-api-gateway \
  --network rvstore \
  --network-alias rvstore-api-gateway \
  -e SPRING_PROFILES_ACTIVE=compose \
  agarwalconsulting/rvstore-api-gateway:v1

docker run --rm -d -p 80:80 \
  --name rvstore-ui \
  --network rvstore \
  --network-alias rvstore-ui \
  agarwalconsulting/rvstore-ui:v1

docker run --rm -d \
  --name rvstore-order-simulator \
  --network rvstore \
  --network-alias rvstore-order-simulator \
  -e SPRING_PROFILES_ACTIVE=compose \
  agarwalconsulting/rvstore-order-simulator:v1
