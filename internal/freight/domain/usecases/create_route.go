package usecases

import (
	"github.com/lucasrgt/go-microservice/internal/freight/domain/entities"
	"github.com/lucasrgt/go-microservice/internal/freight/domain/interfaces"
	"github.com/lucasrgt/go-microservice/internal/freight/infra/repositories"
)

type CreateRouteInput struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Distance float64 `json:"distance"`
	Event    string  `json:"event"`
}

type CreateRouteOutput struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Distance     float64 `json:"distance"`
	Status       string  `json:"status"`
	FreightPrice float64 `json:"freight_price"`
}

type CreateRouteUseCase struct {
	Repository repositories.RouteRepositoryInterface
	Freight    interfaces.FreightInterface
}

func NewCreateRouteUseCase(repository repositories.RouteRepositoryInterface, freight interfaces.FreightInterface) *CreateRouteUseCase {
	return &CreateRouteUseCase{
		Repository: repository,
		Freight:    freight,
	}
}

func (useCase *CreateRouteUseCase) Call(input CreateRouteInput) (*CreateRouteOutput, error) {
	route := entities.NewRoute(input.ID, input.Name, input.Distance)
	useCase.Freight.Calculate(route)
	err := useCase.Repository.Create(route)

	if err != nil {
		return nil, err
	}

	return &CreateRouteOutput{
		ID:           route.ID,
		Name:         route.Name,
		Distance:     route.Distance,
		Status:       route.Status,
		FreightPrice: route.FreightPrice,
	}, nil
}
