PACKAGE_NAME := ysf/learn-grpc

generate:
	protoc -I=./proto --go_out=plugins=grpc:./proto proto/*.proto
