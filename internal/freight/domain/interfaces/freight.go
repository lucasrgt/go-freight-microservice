package interfaces

import "github.com/lucasrgt/go-microservice/internal/freight/domain/entities"

type FreightInterface interface {
	Calculate(route *entities.Route)
}
