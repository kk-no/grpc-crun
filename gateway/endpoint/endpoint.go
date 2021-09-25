package endpoint

import (
	"context"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	idtoken "github.com/salrashid123/oauth2/idtoken"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
)

func Set(ctx context.Context, mux *runtime.ServeMux) error {
	return setSample(ctx, mux)
}

func makeCredentialOptions(ctx context.Context, domain string) ([]grpc.DialOption, error) {
	pool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	cred, err := rpcCredential(ctx, fmt.Sprintf("https://%s", domain))
	if err != nil {
		return nil, err
	}

	return []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(pool, "")),
		grpc.WithPerRPCCredentials(cred),
	}, nil
}

func rpcCredential(ctx context.Context, audience ...string) (credentials.PerRPCCredentials, error) {
	scope := "https://www.googleapis.com/auth/userinfo.email"

	cred, err := google.FindDefaultCredentials(ctx, scope)
	if err != nil {
		return nil, err
	}

	idTokenSource, err := idtoken.IdTokenSource(&idtoken.IdTokenConfig{
		Credentials: cred,
		Audiences:   audience,
	})
	if err != nil {
		return nil, err
	}

	return idtoken.NewIDTokenRPCCredential(ctx, idTokenSource)
}
