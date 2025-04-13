package models

import (
	"net/url"
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
)

type TendersPagination struct {
	Tenders    []Tender
	Pagination pagination.Pagination
}

type Tender struct {
	VerificationObject
	FavouriteObject

	ID                 int
	Name               string
	City               City
	Organization       Organization
	WinnerOrganization Optional[Organization]
	Price              int
	IsContractPrice    bool
	IsNDSPrice         bool
	IsDraft            bool
	FloorSpace         int
	Description        string
	Wishes             string
	Specification      string
	Attachments        []string
	Services           []Service
	Objects            []Object
	VerificationStatus VerificationStatus
	Status             TenderStatus
	ReceptionStart     time.Time
	ReceptionEnd       time.Time
	WorkStart          time.Time
	WorkEnd            time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (t Tender) ToVerificationObject() api.VerificationRequestObject {
	return api.VerificationRequestObject{
		Type:   api.TenderVerificationRequestObject,
		Tender: ConvertTenderModelToApi(t),
	}
}

func (t Tender) ToFavouriteObject() api.FavouritesObject {
	return api.FavouritesObject{
		Type:   api.TenderFavouritesObject,
		Tender: ConvertTenderModelToApi(t),
	}
}

func ConvertTenderModelToApi(tender Tender) api.Tender {
	tenderApi := api.Tender{
		ID:              tender.ID,
		Name:            tender.Name,
		City:            ConvertCityModelToApi(tender.City),
		Organization:    ConvertOrganizationModelToApi(tender.Organization),
		Price:           float64(tender.Price / 100),
		IsContractPrice: tender.IsContractPrice,
		IsNdsPrice:      tender.IsNDSPrice,
		IsDraft:         tender.IsDraft,
		FloorSpace:      tender.FloorSpace,
		Description:     tender.Description,
		Wishes:          tender.Wishes,
		Specification:   stringToUrl(tender.Specification),
		Attachments: convert.Slice[[]string, []url.URL](
			tender.Attachments, func(s string) url.URL { return stringToUrl(s) },
		),
		Services: convert.Slice[[]Service, api.Services](
			tender.Services, ConvertServiceModelToApi,
		),
		Objects: convert.Slice[[]Object, api.Objects](
			tender.Objects, ConvertObjectModelToApi,
		),
		Status:             api.Status(tender.Status.ToStatus()),
		VerificationStatus: api.OptVerificationStatus{Value: tender.VerificationStatus.ToAPI(), Set: tender.VerificationStatus != 0},
		ReceptionStart:     tender.ReceptionStart,
		ReceptionEnd:       tender.ReceptionEnd,
		WorkStart:          tender.WorkStart,
		WorkEnd:            tender.WorkEnd,
		CreatedAt:          tender.CreatedAt,
		UpdatedAt:          tender.UpdatedAt,
	}

	return tenderApi
}

func ConvertServiceModelToApi(service Service) api.Service {
	return api.Service{
		ID:       service.ID,
		ParentID: api.OptInt{Value: service.ParentID, Set: service.ParentID != 0},
		Name:     service.Name,
	}
}

func ConvertObjectModelToApi(object Object) api.Object {
	return api.Object{
		ID:       object.ID,
		ParentID: api.OptInt{Value: object.ParentID, Set: object.ParentID != 0},
		Name:     object.Name,
	}
}

type TenderStatus int

const (
	InvalidStatus TenderStatus = iota
	DraftStatus
	OnModerationStatus
	ReceptionNotStartedStatus
	ReceptionStatus
	SelectingContractorStatus
	ContractorSelectedStatus
	ContractorNotSelectedStatus
	RemovedByModeratorStatus
)

func (s TenderStatus) ToStatus() api.Status {
	switch s {
	case DraftStatus:
		return api.StatusDraft
	case OnModerationStatus:
		return api.StatusOnModeration
	case ReceptionNotStartedStatus:
		return api.StatusReceptionNotStarted
	case ReceptionStatus:
		return api.StatusReception
	case SelectingContractorStatus:
		return api.StatusSelectingContractor
	case ContractorSelectedStatus:
		return api.StatusContractorSelected
	case ContractorNotSelectedStatus:
		return api.StatusContractorNotSelected
	case RemovedByModeratorStatus:
		return api.StatusRemovedFromPublication
	default:
		return api.StatusInvalid
	}
}
