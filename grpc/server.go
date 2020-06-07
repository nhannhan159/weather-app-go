package grpc

import (
	"log"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcOpentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/nhannhan159/weather-app-go/common"
	"github.com/nhannhan159/weather-app-go/domain"
	"google.golang.org/grpc"
)

type GRPCServerManager struct {
	config       common.GRPCConfig
	appResources *domain.Resources
	handlers     []domain.IGRPCHandler
}

func NewGRPCServerManager(config common.GRPCConfig) *GRPCServerManager {
	return &GRPCServerManager{
		config:   config,
		handlers: []domain.IGRPCHandler{},
	}
}

func (server *GRPCServerManager) RegisterResources(resources *domain.Resources) {
	server.appResources = resources
}

func (server *GRPCServerManager) RegisterHandler(handler domain.IGRPCHandler) {
	server.handlers = append(server.handlers, handler)
}

func (server *GRPCServerManager) Run() {
	lis, err := net.Listen("tcp", ":"+server.config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcCtxtags.StreamServerInterceptor(),
			grpcOpentracing.StreamServerInterceptor(),
			grpcZap.StreamServerInterceptor(server.appResources.GRPCLogger),
			grpcRecovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcCtxtags.UnaryServerInterceptor(),
			grpcOpentracing.UnaryServerInterceptor(),
			grpcZap.UnaryServerInterceptor(server.appResources.GRPCLogger),
			grpcRecovery.UnaryServerInterceptor(),
		)),
	)
	for _, handler := range server.handlers {
		handler.RegisterServer(s)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
