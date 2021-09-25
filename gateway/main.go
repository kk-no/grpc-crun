package main

import (
	"context"
	"github.com/kk-no/grpc-crun/gateway/endpoint"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kk-no/grpc-crun/gateway/config"
)

func main() {
	ctx := context.Background()

	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	conf := config.Conf

	mux := runtime.NewServeMux()
	if err := endpoint.Set(ctx, mux); err != nil {
		return err
	}

	log.Println("Gateway server listen on", conf.Port)
	return http.ListenAndServe(":"+conf.Port, mux)
}
