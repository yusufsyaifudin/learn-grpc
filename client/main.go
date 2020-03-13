package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"
	"ysf/learn-grpc/pkg/tracer"
	service "ysf/learn-grpc/proto"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"google.golang.org/grpc"
)

const (
	tracerURL   = "localhost:5775"
	serviceName = "GRPC-SERVICE-EXAMPLE"
)

var (
	serverAddr = flag.String("server_addr", "localhost:3000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	tracerService, closer, err := tracer.New(serviceName, tracerURL)
	if err != nil {
		log.Fatalf("error on tracer: %s", err.Error())
		return
	}

	defer func() {
		_ = closer.Close()
	}()

	var opts = []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(tracerService),
		),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(tracerService),
		),
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return
	}

	defer func() {
		_ = conn.Close()
	}()

	client := service.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call Echo Service on function Echo
	resp, err := client.Echo(ctx, &service.EchoRequest{
		Message: "Hello World!",
	})

	if err != nil {
		log.Fatalf("err making request: %v", err)
		return
	}

	fmt.Println(resp)
}
