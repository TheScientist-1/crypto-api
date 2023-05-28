# Cryptocurrency Rate Notification Service
This project is a solution for fetching real-time cryptocurrency rates and dispatching them to subscribed emails. It has been implemented using a microservice architecture. It consists of four services, with communication between them. 

## Services

* Rate Fetcher Service
* Subscription Service
* Notification Service
* API Gateway

## Docker

Build docker image:

```docker
docker build -t crypto_api .
```

Build docker container:

```docker
docker run -p 8080:8080 crypto_api
```