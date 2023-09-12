1、都在本机可以访问
2、服务端在docker容器中，客户端在主机不能访问
3、服务端和客户端在都在容器环境下能够访问

总结，要在一个局域网下，要能发送广播

```bash
go get go-micro.dev/v4

go mod tidy


GOOS=linux GOARCH=amd64 go build -o ./greeter-srv .

GOOS=linux GOARCH=amd64 go build -o ./greeter-cli .

docker build -t microsrv .

docker build -t microcli .

docker run -it --rm --name microsrv microsrv

docker run -it --rm --name microcli microcli
```

# Greeter

An example Greeter application

## Contents

- **srv** - an RPC greeter service
- **cli** - an RPC client that calls the service once

## Run Service

Start go.micro.srv.greeter
```shell
go run srv/main.go
```

## Run Client

Call go.micro.srv.greeter via client
```shell
go run cli/main.go
```

Examples of client usage via other languages can be found in the client directory.
