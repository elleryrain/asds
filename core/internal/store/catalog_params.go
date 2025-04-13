package store

type CatalogCreateRegionParams struct {
	Name string
}

type CatalogCreateCityParams struct {
	Name     string
	RegionID int
}

type CatalogCreateObjectParams struct {
	Name     string
	ParentID int
}

type CatalogCreateServiceParams struct {
	Name     string
	ParentID int
}
