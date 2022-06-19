# MiG-28

This repo contains a take on My idiomatic Go code.

## cmd/
* This is where you'll find the executables (cli tool etc) and `main` package.
* Ideally, `main` package ties down all the dependencies.

## config/
* Centralized point for application configuration variables.

## core/
* This package holds the main business logic of the application.

## mocks/
* This package holds mock implementations of application behaviour.

## server/
* This package holds implementation related to server (http, grpc etc).

## storage/
* This package holds implementation for storage layer i.e databases.

## root/
* Contains our domain types, Makefile, Dockerfile etc. Files that don't need to be referenced/packaged.
