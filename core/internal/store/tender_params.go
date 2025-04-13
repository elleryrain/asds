package store

import (
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

type TenderCreateParams struct {
	Name            string
	ServiceIDs      []int
	ObjectIDs       []int
	Price           int
	IsContractPrice bool
	IsNDSPrice      bool
	IsDraft         bool
	CityID          int
	FloorSpace      int
	Description     string
	Wishes          string
	Specification   string
	Attachments     []string
	ReceptionStart  time.Time
	ReceptionEnd    time.Time
	WorkStart       time.Time
	WorkEnd         time.Time
	OrganizationID  int
}

type TenderListParams struct {
	OrganizationID models.Optional[int]
	TenderIDs      models.Optional[[]int]
	WithDrafts     bool
	VerifiedOnly   bool
	Limit          models.Optional[uint64]
	Offset         models.Optional[uint64]
}

type TenderUpdateParams struct {
	ID              int
	Name            models.Optional[string]
	ServiceIDs      models.Optional[[]int]
	ObjectIDs       models.Optional[[]int]
	Price           models.Optional[int]
	IsContractPrice models.Optional[bool]
	IsNDSPrice      models.Optional[bool]
	IsDraft         models.Optional[bool]
	CityID          models.Optional[int]
	FloorSpace      models.Optional[int]
	Description     models.Optional[string]
	Wishes          models.Optional[string]
	Specification   models.Optional[string]
	Attachments     models.Optional[[]string]
	ReceptionStart  models.Optional[time.Time]
	ReceptionEnd    models.Optional[time.Time]
	WorkStart       models.Optional[time.Time]
	WorkEnd         models.Optional[time.Time]
}

type TenderServicesCreateParams struct {
	TenderID    int
	ServicesIDs []int
}

type TenderObjectsCreateParams struct {
	TenderID   int
	ObjectsIDs []int
}

type TenderObjectsDeleteParams struct {
	TenderID   int
	ObjectsIDs []int
}

type TenderServicesDeleteParams struct {
	TenderID    int
	ServicesIDs []int
}

type TenderUpdateVerifStatusParams struct {
	TenderID           int
	VerificationStatus models.VerificationStatus
	Status             models.TenderStatus
}

type TenderGetCountParams struct {
	OrganizationID models.Optional[int]
	TenderIDs      models.Optional[[]int]
	WithDrafts     bool
	VerifiedOnly   bool
}

type TenderUpdateStatusParams struct {
	TenderID int
	Status   models.Optional[models.TenderStatus]
}

type TenderNotifyInfoParams struct {
	TenderID         models.Optional[int]
	AdditionID       models.Optional[int]
	QuestionAnswerID models.Optional[int]
}
