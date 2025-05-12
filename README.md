# mysportapp
Small pet project just to learn Go language and number of SE topics

I'm trying to follow Hexagonal Architecture (Port & Adapter pattern) while building this application.

You can find some diagrams built with Draw.io (please take a look into `diagrams` directory) that describes some aspects of this application.

## Pre-Requisites
- `docker` to make local development and testing much simpler
- `docker-compose` to run the whole stack at once on Docker
- `migrate` to run DB migrations ([repo](https://github.com/golang-migrate/migrate))

## Quick start

To start all the things really quickly please consider the next command:
```bash
$ make start
```
