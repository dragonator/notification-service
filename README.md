# rental-service

[![Go Build](https://github.com/dragonator/notification-service/actions/workflows/go.yml/badge.svg)](https://github.com/dragonator/notification-service/actions/workflows/go.yml)

The service exposes a single endpoint for sending a message to multiple channels:

`GET /notify`

    {
        "message": "<text>"
    }


## Run the service

#### Init environment file:

    make init

#### Start Kafka services:

    docker-compose up -d

The following services should be started:
* kafka-ui
* notification-service-kafka-1
* notification-service-zookeeper-1

#### Run the Kafka consumer:

    make notification-consumer-start

#### Run the service:

    make server-start

