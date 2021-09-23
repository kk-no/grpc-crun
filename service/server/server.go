package server

import (
	"context"
	"log"

	sample "github.com/kk-no/proto-terminal/sample/v1"
)

type SampleServer struct {
	sample.UnimplementedSampleServiceServer
}

func (s *SampleServer) Ping(ctx context.Context, req *sample.PingRequest) (*sample.PingResponse, error) {
	log.Println("Called SampleServer.Ping")
	return &sample.PingResponse{}, nil
}
