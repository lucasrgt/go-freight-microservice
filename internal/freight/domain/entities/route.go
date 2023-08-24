package entities

import "time"

type Route struct {
	ID           string
	Name         string
	Distance     float64
	Status       string
	FreightPrice float64
	StartedAt    time.Time
	FinishedAt   time.Time
}

func NewRoute(id string, name string, distance float64) *Route {
	return &Route{
		ID:       id,
		Name:     name,
		Distance: distance,
		Status:   "pending",
	}
}

func (route *Route) Start(startedAt time.Time) {
	route.Status = "started"
	route.StartedAt = startedAt
}

func (route *Route) Finish(finishedAt time.Time) {
	route.Status = "finished"
	route.FinishedAt = finishedAt
}
