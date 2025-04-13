package amoproxy

import (
	"context"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm/dto"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/models"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/store"
)

const (
	CompanyFieldIDPhone              = 612237
	CompanyFieldIDEmail              = 612239
	CompanyFieldIDWeb                = 612241
	CompanyFieldIDAddress            = 612243
	CompanyFieldIDRole               = 662299
	CompanyFieldIDFullName           = 662649
	CompanyFieldIDShortName          = 662653
	CompanyFieldIDINN                = 662655
	CompanyFieldIDKPP                = 662657
	CompanyFieldIDRegisteredAt       = 662659
	CompanyFieldIDRegion             = 662661
	CompanyFieldIDOKVED2             = 662713
	CompanyFieldIDMainActivity       = 662715
	CompanyFieldIDOKPO               = 662817
	CompanyFieldIDOGRN               = 662819
	CompanyFieldIDTaxCode            = 662821
	CompanyFieldIDLegalAddress       = 662919
	CompanyFieldIDCity               = 667231
	CompanyFieldIDWorkType           = 672615
	CompanyFieldIDActivityField      = 672829
	CompanyFieldIDComment            = 774871
	CompanyFieldIDCompanyDescription = 1112751
	CompanyFieldIDBrandName          = 1112801
	CompanyFieldIDPhone1             = 1271395
	CompanyFieldIDPhone2             = 1271397
	CompanyFieldIDPhone3             = 1271399
	CompanyFieldIDPhone4             = 1271401
	CompanyFieldIDPhone5             = 1271403
	CompanyFieldIDPhone6             = 1271405
	CompanyFieldIDPhone7             = 1271455
	CompanyFieldIDPhone8             = 1271457
	CompanyFieldIDPhone9             = 1271459
)

const (
	ContactFieldIDPosition   = 612235
	ContactFieldIDPhone      = 612237
	ContactFieldIDEmail      = 612239
	ContactFieldIDFirstName  = 740801
	ContactFieldIDLastName   = 740803
	ContactFieldIDMiddleName = 740805
	ContactFieldIDTelegram   = 1112803
	ContactFieldIDWA         = 1112805
)

type Service struct {
	amoCRMGateway AmoCRMGateway
	amoStore      AmoStore
	psql          DBTX
}

type AmoCRMGateway interface {
	CreateCompanies(ctx context.Context, request dto.CreateCompaniesRequest) (dto.BasicResponse[dto.CreateCompaniesResponse], error)
	CreateContacts(ctx context.Context, request dto.CreateContactsRequest) (dto.BasicResponse[dto.CreateContactsResponse], error)
	CreateLinks(ctx context.Context, request dto.CreateLinksRequest) (dto.BasicResponse[dto.CreateLinksResponse], error)
	CreateLeads(ctx context.Context, request dto.CreateLeadsRequest) (dto.BasicResponse[dto.CreateLeadsResponse], error)
	CreateNotes(ctx context.Context, request dto.CreateNotesRequest) error
}

type AmoStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, externalID int, amoID int, entity models.Entity) error
	Get(ctx context.Context, qe store.QueryExecutor, externalID int, entity models.Entity) (int, error)
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

func New(amoCRMGateway AmoCRMGateway, amoStore AmoStore, psql DBTX) *Service {
	return &Service{
		amoCRMGateway: amoCRMGateway,
		amoStore:      amoStore,
		psql:          psql,
	}
}
