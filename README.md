# RV Store

A go-kit sample application. It is also used as an example app in a Kubernetes class.

## Setup

## Services

### Auxillary Services

#### Mongodb

This is the stock Mongo database. Find it on Docker Hub and follow the instructions to run it.

* Container port: 27017
* Container name: mongo

#### API Gateway Service

This is a Java Spring Boot application that runs on Java 8. It is a runnable jar file, which contains its own application server.

* Container port: 9000
* Container name: agarwalconsulting/rvstore-api-gateway
* Application code located in web/gateway-service/target/gateway-service.jar
* Command to start the process: java -jar gateway-service.jar
* Environment variable needed: SPRING_PROFILES_ACTIVE=compose

#### UI

This is a static web site made up of images, HTML files, CSS, and Javascript. Pick a web server and serve up the static files located in the Github repo. I suggest Nginx.

* Container port: 80
* Container name: agarwalconsulting/rvstore-ui
* Static files located in web/rvstore/dist/ui

#### Order Simulator

This container submits a fresh order to the order service once a minute. It does not need to expose any ports to receive traffic since it’s a worker process. It is a Java Spring Boot application that contains the application server in it. It runs on Java 8.

* Container name: agarwalconsulting/rvstore-order-simulator
* Application code located in tests/order-simulator/target/order-simulator-0.0.1-
SNAPSHOT.jar
* Command to start the process: java -jar order-simulator-0.0.1-SNAPSHOT.jar
* Environment variable needed: SPRING_PROFILES_ACTIVE=compose
