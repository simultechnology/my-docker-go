# Golang web server on Docker container

## run

```
go run main.go
```

## build a docker image

```
docker build --tag docker-golang-webserver .

```

## run a docker container

``` 
docker run -p 8080:8080 docker-golang-webserver
```

you can see this url.

[http://localhost:8080/](http://localhost:8080/)

## reference

[https://docs.docker.com/language/golang/build-images/](https://docs.docker.com/language/golang/build-images/)