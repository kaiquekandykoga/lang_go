## Steps to create app

```
$ mkdir basic; cd basic
$ go mod init v0
$ touch main.go
```

## Add the gin dependencies

```
$ go get -u github.com/gin-gonic/gin
```

## Run app

```
$ go run main.go
```

## Access app

```
$ curl localhost:8080/status
$ curl localhost:8080/time
```

# Docker

```
nerdctl build -t lang-go-gin-basic .
nerdctl run --rm -p 8080:8080 lang-go-gin-basic
```

