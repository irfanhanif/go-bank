# go-bank-interface

## Description
This is an application interface for A Simple Bank Transaction software. The aim of this project is to learn how tocreate a Go app from scratch.

## Objectives
As the project will grow in the future, this project is designed to be able to communicate with another outside service trough several interfaces such as Web, REST, SOAP, Kafka, RabbitMQ, etc. Therefore, the structure of the project would be critical as the code need to be easy to understand, maintain, and adapt.

## How to Run
Navigate to `cmd` directory. To run the app with Web interface, go to `server` directory. Then execute a command below:
```
go run main.go --host=localhost --port=8080 --intface=web
```