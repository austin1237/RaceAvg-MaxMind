# RaceAvg-MaxMind
[![CircleCI](https://circleci.com/gh/austin1237/RaceAvg-MaxMind.svg?style=svg)](https://circleci.com/gh/austin1237/RaceAvg-MaxMind)
## Prerequisites
You must have the following installed/configured for this to work correctly<br />
1. [Docker](https://www.docker.com/community-edition)
2. [Docker-Compose](https://docs.docker.com/compose/)

## Why Docker and Docker-Compose?
1. Docker is a great way to handle dependencies from the OS to the packages for a specific languages. It's like vagrant but you can deploy it.<br>
2. Docker-Compose is a very nice abstraction layer for running docker commands on your system.

## Race Average

### Development Environment
To spin up the development environment run the following command from the root level of this repo.

```bash
docker-compose -f raceavg.yml up
```

This container uses [fresh](https://github.com/pilu/fresh) as a hot-reloader. That way whenever a source file changes in the raceavg dir the container will automatically rebuild/run the program.

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

This container uses [fresh](https://github.com/pilu/fresh) as a hot-reloader. That way whenever a source file changes in the mmapi dir the container will automatically rebuild/run the program.<br>

### Documentation
runs on http://localhost:3001 </br>
uses http://apidocjs.com/ <br>

### Where do I put the CSV files.
You don't. The build process will download the needed csv files for you, and place them in the container's home/maxmind dir. This removes the need to store large files in git.

### Test Suite
To run the test suite run the following command from the root level of this repo

```bash
docker-compose -f mmapi.yml run mmapi go test -v ./...
```

### Design decisions.
I took the "fit into memory" and no db needed literally and load both csv files into memory on start. This slows down the initial services start time but you gain fast in memory lookups.

The following structures compose the in memory "db" for querying.

1. A sorted list of startIps in ascending order from the blocks csv.
2. Hashmap[startIp]block from the blocks csv.
3. Hashmap[locId]location from the locations csv.

To find the geolocation of an ip I replicated this [query](https://dev.maxmind.com/geoip/legacy/csv/#SQL_Queries).
1. Find the range of viable start ips compared to the client's ip
2. Look through the range and found the bloc from the hashmap[startIp]block that has a endIp that is <= the client ip
3. Return the location with the same bloc.locId form step 2 from the Hashmap[locId]location.


###Why use Go over Node/Python/Java or any other language for the api.
Not considering performance which Go is great at. From an ops perspective since GO creates a static binary you can create very lean/small container's in the CI pipeline that don't even need GO's runtime to run. Also bonus points...

