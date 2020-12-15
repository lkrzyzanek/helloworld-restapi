# helloworld-restapi

Minimalistic hello world REST API written Go.

On any request it returns

```json
{
  "message": "hello world"
}
```

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

