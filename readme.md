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
    --network sundae \
    -p 27017:27017 \
    -e MONGO_INITDB_ROOT_USERNAME=admin \
    -e MONGO_INITDB_ROOT_PASSWORD=pwd \
    -e MONGO_INITDB_DATABASE=sundae \
    -v $PWD/init_mongo_test.js:/docker-entrypoint-initdb.d/init_mongo_test.js \
    mongo
```

Run the tests, if custom network was created for db container add the dev container in this same docker network.

```bash
make go_test
```
