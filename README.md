# LEARN gRPC

Learn how to code gRPC server and client in Golang. 
This is only echoing request message payload into response message payload using Unary call.

[See presentation here.](https://docs.google.com/presentation/d/11nN_yeQEsy-dlrwEEaLIYPcm1y1nqhpcNVzTP12B8Fo/edit?usp=sharing)

## Install and run

```
git clone git@github.com:yusufsyaifudin/learn-grpc.git
go mod download
```

Then run the server first:

```
go run server/main.go
```

After server already running, call gRPC using client:

```
go run client/main.go
```

Or using https://github.com/fullstorydev/grpcurl

First install it, then generate `protoset` to make `grpcurl` know the descriptor:

```
protoc --proto_path=./proto --descriptor_set_out=learngrpc.protoset --include_imports proto/*.proto
```

List the service without `protoset`:

```
grpcurl -v -plaintext -import-path ./proto -proto echo.proto localhost:3000 list
```

Or with `protoset`:

```
grpcurl -protoset learngrpc.protoset localhost:3000 list
```

Then call the service with plaintext mode:

```
grpcurl -v -plaintext -protoset learngrpc.protoset -d '{"message": "Hello World"}' localhost:3000 service.EchoService/Echo
```

Will look like this:

![Image](https://raw.githubusercontent.com/yusufsyaifudin/learn-grpc/master/assets/img/Screen_Shot_2020-03-13_14.05.29.png)

## Generate Proto
Before generate, please install `protoc` command:

http://google.github.io/proto-lens/installing-protoc.html

MacOS:

```
brew install protobuf
```

Linux:

```
PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP
```

Then install plugin for Go, to generate to go file from `.protoc`

https://github.com/golang/protobuf

```
go get -u -v github.com/golang/protobuf/protoc-gen-go
```

Generate protocol buffer file:

```
make generate
```

## Server

See directory `server`

## Client 

See directory `client`