package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/praktikum-examples/calculator/api/v1"
	grpctransport "github.com/soheilhy/grpctransport/go"
	"google.golang.org/grpc"
)

// RunGRPCServer запускает GRPC сервер.
func RunGRPCServer(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("net.Listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, NewCalculatorService())

	log.Printf("Запуск gRPC сервера на адресе %s", addr)
	return s.Serve(lis)
}

// RunHTTPServer запускает HTTP сервер с использованием GRPC Gateway.
func RunHTTPServer(grpcAddr, httpAddr string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw := runtime.NewServeMux()
	opts := []grpctransport.Option{
		grpctransport.WithDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return grpc.DialContext(ctx, addr, grpc.WithInsecure())
		}),
	}
	err := pb.RegisterCalculatorServiceHandlerFromEndpoint(ctx, gw, grpcAddr, opts...)
	if err != nil {
		return fmt.Errorf("RegisterCalculatorServiceHandlerFromEndpoint: %v", err)
	}

	log.Printf("Запуск HTTP сервера на адресе %s", httpAddr)
	return http.ListenAndServe(httpAddr, gw)
}
