# Memcached as a microservice
An OMG service for Memcached, it is a general-purpose distributed memory caching system.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
<!-- [![Build Status](https://travis-ci.com/heaptracetechnology/microservice-firebase.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-firebase)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-firebase/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-firebase)
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-web.svg)](https://golangci.com) -->

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Set Key-value cache
```sh
$ omg run set -a key=<SET_KEY> -a value=<SET_VALUE> -e MEMCACHED_HOST=<HOST_ADDRESS> -e MEMCACHED_PORT=<PORT_NUMBER>
```
##### Get value pair cache
```sh
$ omg run get -a key=<SET_KEY> -e MEMCACHED_HOST=<HOST_ADDRESS> -e MEMCACHED_PORT=<PORT_NUMBER>
```
## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-memcached .
```
### RUN
```
docker run -p 8000:8000 microservice-memcached
```
