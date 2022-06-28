# MiG-28

This repo contains a take on My idiomatic Go code.

Important concepts studied and used:
* SOLID Principles, Clean Code
* Hexagonal Architecture (Ports & Adapters Pattern)
* Separation of Concern (API, Application, Storage Layer)
* Dependency Injection
* Flat-Hierarchy Structure
* Unit Testing (using mocks)
* Persistence Ignorance

## cmd/
* This is where you'll find the executables (cli tool etc.) and `main` package.
* Ideally, `main` package ties down all the dependencies.

## config/
* Centralized point for application configuration variables.

## core/
* This package holds the main business logic of the application. Also termed as application layer.

## mocks/
* This package holds mock implementations of application behaviour.

## server/
* This package holds implementation related to server (http, grpc etc.). Also termed as API layer.

## storage/
* This package holds implementation for storage layer i.e. databases.

## root/
* Contains our domain types, Makefile, Dockerfile etc. Files that don't need to be referenced/packaged/encapsulated.

## Note
This repository is not in complete/finished state (never will be :P). I will keep adding modules and techs as I go to grow the application into something large scale and reflect real-world example all the while preserving the aforementioned principals. Also, this is my personal and subjective take on these principals and it's implementation. Anyone can disagree or better yet, share constructive criticism to improve the code-base and further my/their knowledge. :)
