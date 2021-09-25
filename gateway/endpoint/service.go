package endpoint

import (
	"context"
	"fmt"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kk-no/grpc-crun/gateway/config"
	sample "github.com/kk-no/proto-terminal/sample/v1"
	"google.golang.org/grpc"
)

func setSample(ctx context.Context, mux *runtime.ServeMux) error {
	conf := config.Conf
	endpoint := fmt.Sprintf("%s:%s", conf.ServiceDomain, conf.ServicePort)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	if os.Getenv("K_SERVICE") != "" {
		credOpts, err := makeCredentialOptions(ctx, conf.ServiceDomain)
		if err != nil {
			return err
		}
		opts = credOpts
	}
	return sample.RegisterSampleServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
