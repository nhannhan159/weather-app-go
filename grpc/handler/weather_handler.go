package handler

import (
	"context"

	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/nhannhan159/weather-app-go/domain"
	pb "github.com/nhannhan159/weather-app-go/grpc/pb/weather"
	"github.com/nhannhan159/weather-app-go/model"
)

type IWeatherService interface {
	FindAll(*domain.Resources) ([]model.Weather, error)
	FindByID(*domain.Resources, int) (*model.Weather, error)
}

type WeatherHandler struct {
	pb.UnimplementedWeatherServiceServer
	appResources   *domain.Resources
	weatherService IWeatherService
}

func NewWeatherHandler(appResources *domain.Resources, weatherService IWeatherService) *WeatherHandler {
	return &WeatherHandler{
		appResources:   appResources,
		weatherService: weatherService,
	}
}

func (handler *WeatherHandler) RegisterServer(s *grpc.Server) {
	pb.RegisterWeatherServiceServer(s, handler)
}

func (handler *WeatherHandler) GetWeather(ctx context.Context, req *pb.WeatherRequest) (*pb.Weather, error) {
	res, err := handler.weatherService.FindByID(handler.appResources, int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	weather := &pb.Weather{
		Id:          int64(res.ID),
		Main:        res.Main,
		Description: *res.Description,
		Icon:        res.Icon,
	}
	return weather, nil
}
