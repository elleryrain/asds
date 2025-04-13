package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
)

type OrganizationsPagination struct {
	Organizations []Organization
	Pagination    pagination.Pagination
}

type ContactInfo struct {
	Contact string `json:"contact"`
	Info    string `json:"info"`
}

type ContactInfos []ContactInfo

func (a ContactInfos) Value() (driver.Value, error) {
	if a == nil {
		return []byte("[]"), nil
	}

	return json.Marshal(a)
}

func (a *ContactInfos) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Organization struct {
	VerificationObject
	FavouriteObject

	ID                 int
	BrandName          string
	FullName           string
	ShortName          string
	INN                string
	OKPO               string
	OGRN               string
	KPP                string
	TaxCode            string
	Address            string
	AvatarURL          string
	VerificationStatus VerificationStatus
	IsContractor       bool
	IsBanned           bool
	Emails             ContactInfos
	Phones             ContactInfos
	Messengers         ContactInfos
	CustomerInfo       CustomerInfo
	ContractorInfo     ContractorInfo
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (o Organization) ToVerificationObject() api.VerificationRequestObject {
	return api.VerificationRequestObject{
		Type:         api.OrganizationVerificationRequestObject,
		Organization: ConvertOrganizationModelToApi(o),
	}
}

func ConvertOrganizationModelToApi(organization Organization) api.Organization {
	return api.Organization{
		ID:                 organization.ID,
		BrandName:          api.Name(organization.BrandName),
		FullName:           api.Name(organization.FullName),
		ShortName:          api.Name(organization.ShortName),
		VerificationStatus: api.NewOptVerificationStatus(organization.VerificationStatus.ToAPI()),
		IsContractor:       organization.IsContractor,
		IsBanned:           organization.IsBanned,
		Inn:                api.Inn(organization.INN),
		Okpo:               api.Okpo(organization.OKPO),
		Ogrn:               api.Ogrn(organization.OGRN),
		Kpp:                api.Kpp(organization.KPP),
		TaxCode:            api.TaxCode(organization.TaxCode),
		Address:            organization.Address,
		AvatarURL:          api.OptURI{Value: stringToUrl(organization.AvatarURL), Set: organization.AvatarURL != ""},
		Emails: convert.Slice[ContactInfos, []api.ContactInfo](
			organization.Emails, ConvertContactInfoModelToApi,
		),
		Phones: convert.Slice[ContactInfos, []api.ContactInfo](
			organization.Phones, ConvertContactInfoModelToApi,
		),
		Messengers: convert.Slice[ContactInfos, []api.ContactInfo](
			organization.Messengers, ConvertContactInfoModelToApi,
		),
		CreatedAt: organization.CreatedAt,
		UpdatedAt: organization.UpdatedAt,
	}
}

func ConvertContactInfoModelToApi(info ContactInfo) api.ContactInfo {
	return api.ContactInfo{
		Contact: info.Contact,
		Info:    info.Info,
	}
}

func ConvertAPIToContactInfo(info api.ContactInfo) ContactInfo {
	return ContactInfo{
		Contact: info.Contact,
		Info:    info.Info,
	}
}

type CustomerInfo struct {
	Description Optional[string] `json:"description"`
	CityIDs     []int            `json:"city_ids"`
	Cities      []City           `json:"-"`
}

func ConvertCustomerInfoToApi(info CustomerInfo) api.CustomerInfo {
	return api.CustomerInfo{
		Description: api.OptDescription{Value: api.Description(info.Description.Value), Set: info.Description.Set},
		Cities:      convert.Slice[[]City, []api.City](info.Cities, ConvertCityModelToApi),
	}
}

func (a CustomerInfo) isEmpty() bool {
	return !a.Description.Set &&
		len(a.CityIDs) == 0 &&
		len(a.Cities) == 0
}

func (a CustomerInfo) Value() (driver.Value, error) {
	if a.isEmpty() {
		return "{}", nil
	}
	return json.Marshal(a)
}

func (a *CustomerInfo) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type ContractorInfo struct {
	Description Optional[string]   `json:"description"`
	Services    []ServiceWithPrice `json:"services"`
	Cities      []City             `json:"-"`
	Objects     []Object           `json:"-"`

	CityIDs   []int `json:"city_ids"`
	ObjectIDs []int `json:"object_ids"`
}

func ConvertContractorInfoToApi(info ContractorInfo) api.ContractorInfo {
	return api.ContractorInfo{
		Description: api.OptDescription{Value: api.Description(info.Description.Value), Set: info.Description.Set},
		Cities:      convert.Slice[[]City, []api.City](info.Cities, ConvertCityModelToApi),
		Objects:     convert.Slice[[]Object, []api.Object](info.Objects, ConvertObjectModelToApi),
		Services:    convert.Slice[[]ServiceWithPrice, []api.ServiceWithPrice](info.Services, ConvertServiceWithPriceToAPI),
	}
}

func (a ContractorInfo) isEmpty() bool {
	return !a.Description.Set &&
		len(a.Services) == 0 &&
		len(a.Cities) == 0 &&
		len(a.Objects) == 0 &&
		len(a.CityIDs) == 0 &&
		len(a.ObjectIDs) == 0
}

func (a ContractorInfo) Value() (driver.Value, error) {
	if a.isEmpty() {
		return "{}", nil
	}
	return json.Marshal(a)
}

func (a *ContractorInfo) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func ConvertContractorModelToApi(org Organization) api.V1OrganizationsContractorsGetOKDataItem {
	return api.V1OrganizationsContractorsGetOKDataItem{
		Organization: ConvertOrganizationModelToApi(org),
		Profile:      ConvertContractorInfoToApi(org.ContractorInfo),
	}
}

func (o OrganizationWithProfile) ToFavouriteObject() api.FavouritesObject {
	return api.FavouritesObject{
		Type:                    api.OrganizationWithProfileFavouritesObject,
		OrganizationWithProfile: ConvertOrganizationWithProfileToApi(o),
	}
}

func ConvertOrganizationWithProfileToApi(org OrganizationWithProfile) api.OrganizationWithProfile {
	return api.OrganizationWithProfile{
		Organization: ConvertOrganizationModelToApi(org.Organization),
		Profile:      ConvertContractorInfoToApi(org.Profile),
	}
}

type OrganizationWithProfile struct {
	FavouriteObject

	Organization Organization   `json:"organization"`
	Profile      ContractorInfo `json:"profile"`
}
