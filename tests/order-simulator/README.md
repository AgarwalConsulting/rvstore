# Order Simulator

> Original code: https://github.com/VergeOps/k8s-rvstore/tree/master/Microservices/order-simulator

This container submits a fresh order to the order service once a minute. It does not need to expose any ports to receive traffic since it’s a worker process. It is a Java Spring Boot application that contains the application server in it. It runs on Java 8.
