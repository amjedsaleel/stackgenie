### How to build

```sh
CGO_ENABLED="0" go build -ldflags="-s -w" -o app main.go
```

### How to run locally

```sh
go run main.go
```

### How to run in a server
```sh
./app -listen 0.0.0.0:80
```

