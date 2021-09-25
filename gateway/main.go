package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kk-no/grpc-crun/gateway/config"
	sample "github.com/kk-no/proto-terminal/sample/v1"
	"google.golang.org/grpc"
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
	opts := []grpc.DialOption{grpc.WithInsecure()}

	endpoint := fmt.Sprintf("%s:%s", conf.ServiceDomain, conf.ServicePort)
	if err := sample.RegisterSampleServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}

	log.Println("Gateway server listen on", conf.Port)
	return http.ListenAndServe(":"+conf.Port, mux)
}
