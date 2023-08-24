package usecases

import (
	"github.com/lucasrgt/go-microservice/internal/freight/core/utils"
	"github.com/lucasrgt/go-microservice/internal/freight/infra/repositories"
	"time"
)

type ChangeRouteStatusInput struct {
	ID         string           `json:"id"`
	StartedAt  utils.CustomTime `json:"started_at"`
	FinishedAt utils.CustomTime `json:"finished_at"`
	Event      string           `json:"event"`
}

type ChangeRouteStatusOutput struct {
	ID         string           `json:"id"`
	Status     string           `json:"status"`
	StartedAt  utils.CustomTime `json:"started_at"`
	FinishedAt utils.CustomTime `json:"finished_atf"`
}

type ChangeRouteStatusUseCase struct {
	Repository repositories.RouteRepositoryInterface
}

func NewChangeRouteStatusUseCase(repository repositories.RouteRepositoryInterface) *ChangeRouteStatusUseCase {
	return &ChangeRouteStatusUseCase{
		Repository: repository,
	}
}

func (useCase *ChangeRouteStatusUseCase) Call(input ChangeRouteStatusInput) (*ChangeRouteStatusOutput, error) {
	route, err := useCase.Repository.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	if input.Event == "RouteStarted" {
		route.Start(time.Time(input.StartedAt))
	}

	if input.Event == "RouteFinished" {
		route.Finish(time.Time(input.FinishedAt))
	}

	err = useCase.Repository.Update(route)
	if err != nil {
		return nil, err
	}

	return &ChangeRouteStatusOutput{
		ID:         route.ID,
		Status:     route.Status,
		StartedAt:  utils.CustomTime(route.StartedAt),
		FinishedAt: utils.CustomTime(route.FinishedAt),
	}, nil
}
