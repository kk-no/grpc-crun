package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	sample "github.com/kk-no/proto-terminal/sample/v1"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := sample.RegisterSampleServiceHandlerFromEndpoint(ctx, mux, "localhost:8080", opts); err != nil {
		return err
	}

	log.Println("Gateway server listen on", port)
	return http.ListenAndServe(":"+port, mux)
}
