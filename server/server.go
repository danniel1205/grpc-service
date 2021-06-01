package server

import (
	"context"
	"fmt"

	"github.com/danniel1205/grpc-service/helloservice"
)

type grpcServer struct {
	helloservice.UnimplementedHelloServiceServer
}

func NewServer() *grpcServer {
	return &grpcServer{}
}

func (s *grpcServer) SayHello(ctx context.Context, req *helloservice.Request) (*helloservice.Response, error) {
	fmt.Println(fmt.Sprintf("SayHello() is invoked by user %s who is from %s", req.GetName(), req.GetFrom()))

	name := req.GetName()
	from := req.GetFrom()

	var resp *helloservice.Response
	if name != "" && from != "" {
		resp = &helloservice.Response{
			Message: fmt.Sprintf("Hello %s, %s is a nice place", name, from),
		}
	} else {
		resp = &helloservice.Response{
			Message: fmt.Sprint("Hello"),
		}
	}

	return resp, nil
}
