package service

type RespondCreateParams struct {
	TenderID       int
	OrganizationID int
	Price          int
	IsNds          bool
}

type RespondGetParams struct {
	TenderID int
	Page     uint64
	PerPage  uint64
}
