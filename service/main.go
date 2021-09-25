package main

import (
	"log"
	"net"

	"github.com/kk-no/grpc-crun/service/config"
	"github.com/kk-no/grpc-crun/service/server"
	sample "github.com/kk-no/proto-terminal/sample/v1"
	"google.golang.org/grpc"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	conf := config.Conf

	grpcServer := grpc.NewServer()
	sample.RegisterSampleServiceServer(grpcServer, &server.SampleServer{})

	log.Println("Service server listen on", conf.Port)
	listenPort, err := net.Listen("tcp", ":"+conf.Port)
	if err != nil {
		log.Fatalln(err)
	}
	return grpcServer.Serve(listenPort)
}
