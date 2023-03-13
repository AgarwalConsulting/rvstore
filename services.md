# Services

There are six services plus a Mongo DB, each with their own Docker image:

- Angular UI running in Nginx
- Product service
- Order service
- Auth service
- Order simulator
- Gateway edge service

## Backend Services

### `products` API

This is a Go application. It serves up the product information as a REST API.

- Docker image: `agarwalconsulting/rvstore-product-api`
- Container port: `9001`
- The application should only be accessible inside the cluster

### `order` API

This is a Go application. It receives order data and stores it in the Mongo database.

- Docker image: `agarwalconsulting/rvstore-order-api`
- Container port: `9002`
- The application should only be accessible inside the cluster
- It communicates with the Mongo connection string provided by: `MONGO_DB_URL`

### `auth` API

This is a Go application. It provides authentication, via JWT tokens.

- Docker image: `agarwalconsulting/rvstore-auth-api`
- Container port: `9003`
- The application should only be accessible inside the cluster

## Auxillary Services

### Mongodb

This is the stock Mongo database. Find it on Docker Hub and follow the instructions to run it.

- Docker image: `mongo`
- Container port: `27017`
- Should be accessible only within the cluster
- Mongo stores data internally at `/data/db`
- Environment variables needed:
  - `MONGO_INITDB_ROOT_USERNAME: mongoadmin`
  - `MONGO_INITDB_ROOT_PASSWORD: secret`

### API Gateway Service

This is a nginx proxy server.

- Docker image: `agarwalconsulting/rvstore-api-gateway`
- Container port: `9000`
- Application code located under `web/gateway-service/`

### UI

This is a static web site made up of images, HTML files, CSS, and Javascript. Pick a web server and serve up the static files located in the Github repo. I suggest Nginx.

- Docker image: `agarwalconsulting/rvstore-ui`
- Container port: `80`
- Static files located in `web/rvstore/dist/ui`

### Order Simulator

This is a Java application. It continually submits a fresh order to the order service once a minute. It does not need to expose any ports to receive traffic since it’s a worker process. It runs on Java 8.

- Docker image: `agarwalconsulting/rvstore-order-simulator`
- Application code located in `tests/order-simulator/`
- Environment variable needed: `SPRING_PROFILES_ACTIVE=compose`
