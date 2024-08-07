# go-example-backend 2

This project demonstrates a Go server, that writes to and reads from a Redis database.

## Running Go tests locally

Download and run the `redis` docker image:

```
docker pull redis
docker run --name redis-test-instance -p 6379:6379 -d redis
```

Run Go tests:

```
go test ./...
```
