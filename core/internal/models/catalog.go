package models

import api "gitlab.ubrato.ru/ubrato/core/api/gen"

type Service struct {
	ID       int
	ParentID int
	Name     string
}

type Services []Service

func ConvertModelServiceToApi(c Service) api.Service {
	return api.Service{
		ID:       c.ID,
		ParentID: api.OptInt{Value: c.ParentID, Set: c.ParentID != 0},
		Name:     c.Name,
	}
}

type Object struct {
	ID       int
	ParentID int
	Name     string
}

type Objects []Object

func ConvertModelObjectToApi(c Object) api.Object {
	return api.Object{
		ID:       c.ID,
		ParentID: api.OptInt{Value: c.ParentID, Set: c.ParentID != 0},
		Name:     c.Name,
	}
}

type Measure struct {
	ID   int
	Name string
}

func ConvertMeasureToAPI(m Measure) api.Measure {
	return api.Measure{
		ID:   m.ID,
		Name: m.Name,
	}
}

type ServiceWithPrice struct {
	ServiceID int     `json:"service_id"`
	MeasureID int     `json:"measure_id"`
	Price     float32 `json:"price"`

	Service Service `json:"-"`
	Measure Measure `json:"-"`
}

func ConvertServiceWithPriceToAPI(s ServiceWithPrice) api.ServiceWithPrice {
	return api.ServiceWithPrice{
		Service: ConvertServiceModelToApi(s.Service),
		Measure: ConvertMeasureToAPI(s.Measure),
		Price:   s.Price,
	}
}
