# helloworld-restapi

Minimalistic hello world REST API written Go.

On any request it returns

```json
{
  "message": "hello world",
  "version": 0
}
```

## How to run

```shell
docker run --rm -p 8080:8080 helloworld-restapi:latest
```
you can change `latest` to `1.0` or `2.0` to simulate different APIs.


## Run application locally

Build app

```shell
go build
```

Run

```shell
./helloworld-restapi
```

## Run as local docker image

```shell
docker build . -t helloworld-restapi:latest
```

To run the image execute

```
docker run --rm -p 8080:8080 helloworld-restapi:latest
```

To the image try this URLs
* http://localhost:8080/
* http://localhost:8080/test

## Run on k8s

```shell
kubectl create namespace restapi1
kubectl -n restapi1 apply -f k8s/deployment.yaml
```
