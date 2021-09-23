package main

import (
	"log"
	"net"
	"os"

	"github.com/kk-no/grpc-crun/service/server"
	sample "github.com/kk-no/proto-terminal/sample/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	grpcServer := grpc.NewServer()
	sample.RegisterSampleServiceServer(grpcServer, &server.SampleServer{})

	log.Println("Service server listen on", port)
	listenPort, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	reflection.Register(grpcServer)
	return grpcServer.Serve(listenPort)
}
