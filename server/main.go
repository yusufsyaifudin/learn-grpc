package main

import (
	"context"
	"log"
	"net"
	"ysf/learn-grpc/pkg/tracer"
	service "ysf/learn-grpc/proto"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

const (
	tracerURL   = "localhost:5775"
	serviceName = "GRPC-SERVICE-EXAMPLE"
	port        = ":3000"
)

// server is used to implement service.EchoServiceServer
type echoService struct {
}

func (e echoService) Echo(ctx context.Context, req *service.EchoRequest) (*service.EchoResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Echo")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	return &service.EchoResponse{
		Message: req.Message,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	tracerService, closer, err := tracer.New(serviceName, tracerURL)
	if err != nil {
		log.Fatalf("error on tracer: %s", err.Error())
		return
	}

	defer func() {
		_ = closer.Close()
	}()

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracerService, otgrpc.LogPayloads()),
		),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracerService, otgrpc.LogPayloads()),
		),
	}

	s := grpc.NewServer(opts...)
	service.RegisterEchoServiceServer(s, &echoService{})

	log.Printf("Listening gRPC server at: %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
