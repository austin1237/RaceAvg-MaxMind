# RaceAvg-MaxMind
[![CircleCI](https://circleci.com/gh/austin1237/RaceAvg-MaxMind.svg?style=svg)](https://circleci.com/gh/austin1237/RaceAvg-MaxMind)
## Prerequistes
You must have the following installed/configured for this to work correctly<br />
1. [Docker](https://www.docker.com/community-edition)
2. [Docker-Compose](https://docs.docker.com/compose/)

## Race Average

### Development Environment
To spin up the development environment run the following command from the root level of this repo. 

```bash
docker-compose -f raceavg.yml up
```

This container uses [fresh](https://github.com/pilu/fresh) as a hot-reloader. That way whenever a source file changes in the raceavg dir the container will automacially rebuild/run the program.

### Test Suite
To run the test suite run the following command from the root level of this repo

```bash
docker-compose -f raceavg.yml run raceavg go test -v ./...
```

## MaxMind GeoLite API

### Development Environment
To spin up the development environment run the following command from the root level of this repo. 

```bash
docker-compose -f mmapi.yml up
```

You can communicate with the api over localhost:3000. There is port mapping between the host machine's port 3000 to the container's 8080 port.

This container uses [fresh](https://github.com/pilu/fresh) as a hot-reloader. That way whenever a source file changes in the mmapi dir the container will automacially rebuild/run the program.<br>

### Where do I put the CSV files.
You don't. The build process will download the needed csv files for you, and place them in the container's home/maxmind dir. This removes the need to store large files in git.

### Test Suite
To run the test suite run the following command from the root level of this repo

```bash
docker-compose -f raceavg.yml run raceavg go test -v ./...
```