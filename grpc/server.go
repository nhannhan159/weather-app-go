package grpc

import (
	"log"
	"net"

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

	s := grpc.NewServer()
	for _, handler := range server.handlers {
		handler.RegisterServer(s)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
