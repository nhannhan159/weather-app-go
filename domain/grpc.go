package domain

import (
	"google.golang.org/grpc"
)

type IGRPCManager interface {
	RegisterHandler(IGRPCHandler)
	RegisterResources(*Resources)
	Run()
}

type IGRPCHandler interface {
	RegisterServer(s *grpc.Server)
}
