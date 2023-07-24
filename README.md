# notification-service

[![Go Build](https://github.com/dragonator/notification-service/actions/workflows/go.yml/badge.svg)](https://github.com/dragonator/notification-service/actions/workflows/go.yml)

This is an interview assignment implementation of a notification system which task is to accept user requests and send notification messages on multiple channels. To learn more about the system click [here](/docs/ABOUT.md).

## Run the service

#### Init environment file:

    make init

#### Start Kafka services:

    docker-compose up -d

The following services should be started:
* kafka-ui
* notification-service-kafka-1
* notification-service-zookeeper-1

The server is configured to run on port `9090`.
The command above also starts a `kafka-ui` service that is located on port `8080`.

#### Run the Kafka consumer:

    make notification-consumer-start

#### Run the service:

    make server-start

