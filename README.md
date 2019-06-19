# _Memcached_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-memcached.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-memcached)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-memcached/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-memcached)

An OMG service for Memcached, it is a general-purpose distributed memory caching system.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Set Key-value cache
```coffee
>>> memcached set key:'foo' value:'bar'
{"success":"true/false","message":"success/failure message","statusCode":"HTTPstatusCode"}
```
##### Get value pair cache
```coffee
>>> memcached get key:'foo'
{"key": "foo","value": "bar","expiration": 0,"statuscode": "HTTPstatusCode"}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Set Key-value cache
```shell
$ omg run set -a key=<SET_KEY> -a value=<SET_VALUE> -e MEMCACHED_HOST=<HOST_ADDRESS> -e MEMCACHED_PORT=<PORT_NUMBER>
```
##### Get value pair cache
```shell
$ omg run get -a key=<SET_KEY> -e MEMCACHED_HOST=<HOST_ADDRESS> -e MEMCACHED_PORT=<PORT_NUMBER>
```
**Note**: Start the memcached local server with below command.
```sh
$ docker run -p 11211:11211 --name my-memcache -d memcached memcached -m 64
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/memcached/blob/master/LICENSE).
