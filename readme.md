# api-server

[![Build Status](https://sundae-drone.connan.pro/api/badges/sundae-party/api-server/status.svg)](https://sundae-drone.connan.pro/sundae-party/api-server)

# Build

## Build all the api-server

```bash
make
```

## Test

The default docker bridge network never supported service discovery through a built in DNS.
Create new custom network to enable that, or address db container with container ip ...

```docker
docker network create sundae
```

Create new mongo for the tests:

```docker
docker run \
    --rm -it --name mongo-test \
    -p 27018:27017 \
    -e MONGO_INITDB_ROOT_USERNAME=admin \
    -e MONGO_INITDB_ROOT_PASSWORD=pwd \
    -e MONGO_INITDB_DATABASE=sundae \
    -v $PWD/init_mongo_test.js:/docker-entrypoint-initdb.d/init_mongo_test.js \
    mongo --replSet rs0
```

Run the tests, if custom network was created for db container add the dev container in this same docker network.

```bash
make go_test
```

## API test

curl create / update integration

```bash
curl -X POST -H "Content-Type: application/json" \
    --cacert ca.pem \
    --key cli.key \
    --cert cli.pem \
    -d '{"name": "Hue", "documentation": "https://sundae/doc/hue", "version": "v1.0.1", "url": "https://github.com/sundae-party/integration/hue"}' \
    https://localhost/api/integration
```

```bash
curl -X POST -H "Content-Type: application/json" \
    --cacert ca.pem \
    --key cli.key \
    --cert cli.pem \
    -d '{ "name" : "corridor-sdb_dimmer", "integration" : { "name" : "MQTT" }, "desiredstate" : { "brightness" : 0} }'\
    https://localhost/api/light/desiredstate
```
