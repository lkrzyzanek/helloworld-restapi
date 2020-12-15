# helloworld-restapi

Minimalistic hello world REST API written Go.

On any request it returns

```json
{
  "message": "hello world",
  "version": 1.0
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
docker build . -t helloworld-restapi:1.0
```

To run the image execute

```
docker run --rm -p 8080:8080 helloworld-restapi:1.0
```

To the image try this URLs
* http://localhost:8080/
* http://localhost:8080/test

## Run on Kubernetes

predefined deployment.yaml
```shell
kubectl create namespace restapi
kubectl -n restapi apply -f k8s/deployment.yaml
```

or manually
```shell
kubectl -n restapi create deployment restapi --image lkrzyzanek/helloworld-restapi:latest
kubectl -n restapi expose deployment restapi --type=NodePort --port=8080
```
