package service

import (
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

type TenderListParams struct {
	OrganizationID models.Optional[int]
	WithDrafts     bool
	VerifiedOnly   bool
	Page           uint64
	PerPage        uint64
}

type TenderCreateParams struct {
	Name            string
	OrganizationID  int
	CityID          int
	Price           int
	IsContractPrice bool
	IsNDSPrice      bool
	IsDraft         bool
	FloorSpace      int
	Description     string
	Wishes          string
	Specification   string
	Attachments     []string
	ServiceIDs      []int
	ObjectIDs       []int
	ReceptionStart  time.Time
	ReceptionEnd    time.Time
	WorkStart       time.Time
	WorkEnd         time.Time
}

type TenderUpdateParams struct {
	ID              int
	OrganizationID  int
	Name            models.Optional[string]
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
	ServiceIDs      models.Optional[[]int]
	ObjectIDs       models.Optional[[]int]
	ReceptionStart  models.Optional[time.Time]
	ReceptionEnd    models.Optional[time.Time]
	WorkStart       models.Optional[time.Time]
	WorkEnd         models.Optional[time.Time]
}

type TenderRespondParams struct {
	TenderID       int
	OrganizationID int
	Price          int
	IsNds          bool
}
