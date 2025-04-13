package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type OrganizationGetParams struct {
	IsContractor    models.Optional[bool]
	OrganizationIDs []int
	Offset          models.Optional[uint64]
	Limit           models.Optional[uint64]
}

type OrganizationContractorsGetParams struct {
	Offset models.Optional[uint64]
	Limit  models.Optional[uint64]
}

type OrganizationGetCountParams struct {
	IsContractor    models.Optional[bool]
	OrganizationIDs []int
}

type OrganizationCreateParams struct {
	BrandName    string
	FullName     string
	ShortName    string
	IsContractor bool
	INN          string
	OKPO         string
	OGRN         string
	KPP          string
	TaxCode      string
	Address      string
}

type OrganizationUpdateParams struct {
	OrganizationID int
	Brand          models.Optional[string]
	AvatarURL      models.Optional[string]
	Emails         models.Optional[models.ContactInfos]
	Phones         models.Optional[models.ContactInfos]
	Messengers     models.Optional[models.ContactInfos]
	CustomerInfo   models.Optional[models.CustomerInfo]
	ContractorInfo models.Optional[models.ContractorInfo]
}

type OrganizationUpdateCustomerParams struct {
	OrganizationID int
	CustomerInfo   models.CustomerInfo
}

type OrganizationAddUserParams struct {
	OrganizationID int
	UserID         int
	IsOwner        bool
}

type OrganizationUpdateVerifStatusParams struct {
	OrganizationID     int
	VerificationStatus models.VerificationStatus
}
