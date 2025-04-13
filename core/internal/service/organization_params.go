package service

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type OrganizationGetParams struct {
	IsContractor models.Optional[bool]
	Page         uint64
	PerPage      uint64
}

type OrganizationContractorsGetParams struct {
	Page    uint64
	PerPage uint64
}

type OrganizationUpdateBrandParams struct {
	OrganizationID int
	Brand          models.Optional[string]
	AvatarURL      models.Optional[string]
}

type OrganizationUpdateContactsParams struct {
	OrganizationID int
	Emails         models.Optional[models.ContactInfos]
	Phones         models.Optional[models.ContactInfos]
	Messengers     models.Optional[models.ContactInfos]
}

type OrganizationUpdateContractorParams struct {
	OrganizationID int
	Description    models.Optional[string]
	Services       []models.ServiceWithPrice
	CityIDs        []int
	ObjectIDs      []int
}

type OrganizationUpdateCustomerParams struct {
	OrganizationID int
	Description    models.Optional[string]
	CityIDs        []int
}
