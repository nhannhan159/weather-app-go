package handler

import (
	"context"

	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/nhannhan159/weather-app-go/domain"
	pb "github.com/nhannhan159/weather-app-go/grpc/pb/city"
	"github.com/nhannhan159/weather-app-go/model"
)

type ICityService interface {
	FindAll(*domain.Resources) ([]model.City, error)
	FindByID(*domain.Resources, int) (*model.City, error)
}

type CityHandler struct {
	pb.UnimplementedCityServiceServer
	appResources *domain.Resources
	cityService  ICityService
}

func NewCityHandler(appResources *domain.Resources, cityService ICityService) *CityHandler {
	return &CityHandler{
		appResources: appResources,
		cityService:  cityService,
	}
}

func (handler *CityHandler) RegisterServer(s *grpc.Server) {
	pb.RegisterCityServiceServer(s, handler)
}

func (handler *CityHandler) GetCity(ctx context.Context, req *pb.CityRequest) (*pb.City, error) {
	res, err := handler.cityService.FindByID(handler.appResources, int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	city := &pb.City{
		Id:       int64(res.ID),
		Name:     res.Name,
		FindName: *res.FindName,
		Country:  *res.Country,
		Lat:      res.Lat,
		Lon:      res.Lon,
	}
	return city, nil
}
