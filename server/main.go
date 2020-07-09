package main

import (
	"context"

	"go-grpc-protobuf/proto"
)

type server struct{}

func main() {

}

func (S *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}
