package repositories

import (
	"database/sql"
	"github.com/lucasrgt/go-microservice/internal/freight/domain/entities"
)

type RouteRepositoryInterface interface {
	Create(route *entities.Route) error
	FindByID(id string) (*entities.Route, error)
	Update(route *entities.Route) error
}

type RouteRepositoryMySql struct {
	db *sql.DB
}

func NewRouteRepositorySql(db *sql.DB) *RouteRepositoryMySql {
	return &RouteRepositoryMySql{
		db: db,
	}
}

func (repository *RouteRepositoryMySql) Create(route *entities.Route) error {
	stm := "INSERT INTO routes (id, name, distance, status, freight_price) VALUES (?, ?, ?, ?, ?)"

	_, err := repository.db.Exec(stm, route.ID, route.Name, route.Distance, route.Status, route.FreightPrice)

	if err != nil {
		return err
	}

	return nil
}

func (repository *RouteRepositoryMySql) FindByID(id string) (*entities.Route, error) {
	stm := "SELECT id, name, distance, status, freight_price, started_at, finished_at FROM routes WHERE id = ?"

	row := repository.db.QueryRow(stm, id)

	var startedAt, finishedAt sql.NullTime
	var route entities.Route

	err := row.Scan(
		&route.ID,
		&route.Name,
		&route.Distance,
		&route.Status,
		&route.FreightPrice,
		startedAt,
		finishedAt,
	)

	if err != nil {
		return nil, err
	}

	if startedAt.Valid {
		route.StartedAt = startedAt.Time
	}

	if finishedAt.Valid {
		route.FinishedAt = finishedAt.Time
	}

	return &route, nil
}

func (repository *RouteRepositoryMySql) Update(route *entities.Route) error {
	startedAt := route.StartedAt.Format("2006-01-02 15:04:05")
	finishedAt := route.FinishedAt.Format("2006-01-02 15:04:05")

	stm := "UPDATE routes SET status = ?, freight_price = ?, started_at = ?, finished_at = ? WHERE id = ?"

	_, err := repository.db.Exec(stm, route.Status, route.FreightPrice, startedAt, finishedAt, route.ID)

	if err != nil {
		return err
	}

	return nil
}
