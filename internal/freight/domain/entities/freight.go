package entities

type Freight struct {
	PricePerKm float64
}

func NewFreight(pricePerKm float64) *Freight {
	return &Freight{
		PricePerKm: pricePerKm,
	}
}

func (freight *Freight) Calculate(route *Route) {
	route.FreightPrice = route.Distance * freight.PricePerKm
}
