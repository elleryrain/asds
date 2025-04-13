package models

import api "gitlab.ubrato.ru/ubrato/core/api/gen"

type City struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region Region `json:"region"`
}

func ConvertCityModelToApi(city City) api.City {
	return api.City{
		ID:   city.ID,
		Name: city.Name,
		Region: api.OptRegion{
			Value: ConvertRegionModelToApi(city.Region),
			Set:   city.Region.ID != 0,
		},
	}
}

type Region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ConvertRegionModelToApi(region Region) api.Region {
	return api.Region{
		ID:   region.ID,
		Name: region.Name,
	}
}
