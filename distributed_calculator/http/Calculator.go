package http

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/praktikum-examples/calculator/api/v1"
	"google.golang.org/grpc"
)

// RunGRPCGateway запускает GRPC Gateway.
func RunGRPCGateway(grpcAddr, httpAddr string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterCalculatorServiceHandlerFromEndpoint(ctx, gw, grpcAddr, opts)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", gw)
	return http.ListenAndServe(httpAddr, mux)
}
