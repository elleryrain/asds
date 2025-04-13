package http

import (
	"context"
	"net/http"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

var _ api.Handler = new(Router)

type Router struct {
	Auth
	Error
	Tenders
	Catalog
	Users
	Survey
	Organization
	Suggest
	Verification
	Employee
	Questionnaire
}

type Error interface {
	HandleError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error)
}

type Auth interface {
	V1AuthSigninPost(ctx context.Context, req *api.V1AuthSigninPostReq) (api.V1AuthSigninPostRes, error)
	V1AuthSignupPost(ctx context.Context, req *api.V1AuthSignupPostReq) (api.V1AuthSignupPostRes, error)
	V1AuthUserGet(ctx context.Context) (api.V1AuthUserGetRes, error)
	V1AuthRefreshPost(ctx context.Context, params api.V1AuthRefreshPostParams) (api.V1AuthRefreshPostRes, error)
	V1AuthLogoutPost(ctx context.Context, params api.V1AuthLogoutPostParams) (api.V1AuthLogoutPostRes, error)

	HandleCookieAuth(ctx context.Context, operationName string, t api.CookieAuth) (context.Context, error)
	HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error)
}

type Tenders interface {
	V1TendersPost(ctx context.Context, req *api.V1TendersPostReq) (api.V1TendersPostRes, error)
	V1TendersTenderIDPut(ctx context.Context, req *api.V1TendersTenderIDPutReq, params api.V1TendersTenderIDPutParams) (api.V1TendersTenderIDPutRes, error)
	V1TendersTenderIDGet(ctx context.Context, params api.V1TendersTenderIDGetParams) (api.V1TendersTenderIDGetRes, error)
	V1TendersGet(ctx context.Context, params api.V1TendersGetParams) (api.V1TendersGetRes, error)
	V1OrganizationsOrganizationIDTendersGet(ctx context.Context, params api.V1OrganizationsOrganizationIDTendersGetParams) (api.V1OrganizationsOrganizationIDTendersGetRes, error)

	V1TendersTenderIDAdditionsPost(ctx context.Context, req *api.V1TendersTenderIDAdditionsPostReq, params api.V1TendersTenderIDAdditionsPostParams) (api.V1TendersTenderIDAdditionsPostRes, error)
	V1TendersTenderIDAdditionsGet(ctx context.Context, params api.V1TendersTenderIDAdditionsGetParams) (api.V1TendersTenderIDAdditionsGetRes, error)

	V1TendersTenderIDRespondPost(ctx context.Context, req *api.V1TendersTenderIDRespondPostReq, params api.V1TendersTenderIDRespondPostParams) (api.V1TendersTenderIDRespondPostRes, error)
	V1TendersTenderIDRespondGet(ctx context.Context, params api.V1TendersTenderIDRespondGetParams) (api.V1TendersTenderIDRespondGetRes, error)

	V1TendersTenderIDQuestionAnswerPost(ctx context.Context, req *api.V1TendersTenderIDQuestionAnswerPostReq, params api.V1TendersTenderIDQuestionAnswerPostParams) (api.V1TendersTenderIDQuestionAnswerPostRes, error)
	V1TendersTenderIDQuestionAnswerGet(ctx context.Context, params api.V1TendersTenderIDQuestionAnswerGetParams) (api.V1TendersTenderIDQuestionAnswerGetRes, error)

	V1TendersTenderIDWinnersPost(ctx context.Context, params api.V1TendersTenderIDWinnersPostParams) (api.V1TendersTenderIDWinnersPostRes, error)
	V1TendersTenderIDWinnersGet(ctx context.Context, params api.V1TendersTenderIDWinnersGetParams) (api.V1TendersTenderIDWinnersGetRes, error)
	V1TendersWinnersWinnerIDAprovePost(ctx context.Context, params api.V1TendersWinnersWinnerIDAprovePostParams) (api.V1TendersWinnersWinnerIDAprovePostRes, error)
	V1TendersWinnersWinnerIDDenyPost(ctx context.Context, params api.V1TendersWinnersWinnerIDDenyPostParams) (api.V1TendersWinnersWinnerIDDenyPostRes, error)
}

type Users interface {
	V1UsersUserIDGet(ctx context.Context, params api.V1UsersUserIDGetParams) (api.V1UsersUserIDGetRes, error)
	V1UsersConfirmEmailPost(ctx context.Context, req *api.V1UsersConfirmEmailPostReq) (api.V1UsersConfirmEmailPostRes, error)
	V1UsersConfirmPasswordPost(ctx context.Context, req *api.V1UsersConfirmPasswordPostReq) (api.V1UsersConfirmPasswordPostRes, error)
	V1UsersGet(ctx context.Context, params api.V1UsersGetParams) (api.V1UsersGetRes, error)
	V1UsersUserIDPut(ctx context.Context, req *api.V1UsersUserIDPutReq, params api.V1UsersUserIDPutParams) (api.V1UsersUserIDPutRes, error)
	V1UsersRequestEmailVerificationPost(ctx context.Context, req *api.V1UsersRequestEmailVerificationPostReq) (api.V1UsersRequestEmailVerificationPostRes, error)
	V1UsersRequestResetPasswordPost(ctx context.Context, req *api.V1UsersRequestResetPasswordPostReq) (api.V1UsersRequestResetPasswordPostRes, error)
}

type Survey interface {
	V1SurveyPost(ctx context.Context, req *api.V1SurveyPostReq) (api.V1SurveyPostRes, error)
}

type Catalog interface {
	V1CatalogObjectsGet(ctx context.Context) (api.V1CatalogObjectsGetRes, error)
	V1CatalogServicesGet(ctx context.Context) (api.V1CatalogServicesGetRes, error)
	V1CatalogCitiesPost(ctx context.Context, req *api.V1CatalogCitiesPostReq) (api.V1CatalogCitiesPostRes, error)
	V1CatalogRegionsPost(ctx context.Context, req *api.V1CatalogRegionsPostReq) (api.V1CatalogRegionsPostRes, error)
	V1CatalogObjectsPost(ctx context.Context, req *api.V1CatalogObjectsPostReq) (api.V1CatalogObjectsPostRes, error)
	V1CatalogServicesPost(ctx context.Context, req *api.V1CatalogServicesPostReq) (api.V1CatalogServicesPostRes, error)
	V1CatalogMeasurementsGet(ctx context.Context) (api.V1CatalogMeasurementsGetRes, error)
}

type Organization interface {
	V1OrganizationsOrganizationIDGet(ctx context.Context, params api.V1OrganizationsOrganizationIDGetParams) (api.V1OrganizationsOrganizationIDGetRes, error)
	V1OrganizationsGet(ctx context.Context, params api.V1OrganizationsGetParams) (api.V1OrganizationsGetRes, error)
	V1OrganizationsContractorsGet(ctx context.Context, params api.V1OrganizationsContractorsGetParams) (api.V1OrganizationsContractorsGetRes, error)

	V1OrganizationsOrganizationIDProfileBrandPut(ctx context.Context, req *api.V1OrganizationsOrganizationIDProfileBrandPutReq, params api.V1OrganizationsOrganizationIDProfileBrandPutParams) (api.V1OrganizationsOrganizationIDProfileBrandPutRes, error)
	V1OrganizationsOrganizationIDProfileContactsPut(ctx context.Context, req *api.V1OrganizationsOrganizationIDProfileContactsPutReq, params api.V1OrganizationsOrganizationIDProfileContactsPutParams) (api.V1OrganizationsOrganizationIDProfileContactsPutRes, error)
	V1OrganizationsOrganizationIDProfileContractorPut(ctx context.Context, req *api.V1OrganizationsOrganizationIDProfileContractorPutReq, params api.V1OrganizationsOrganizationIDProfileContractorPutParams) (api.V1OrganizationsOrganizationIDProfileContractorPutRes, error)
	V1OrganizationsOrganizationIDProfileCustomerPut(ctx context.Context, req *api.V1OrganizationsOrganizationIDProfileCustomerPutReq, params api.V1OrganizationsOrganizationIDProfileCustomerPutParams) (api.V1OrganizationsOrganizationIDProfileCustomerPutRes, error)
	V1OrganizationsOrganizationIDProfileContractorGet(ctx context.Context, params api.V1OrganizationsOrganizationIDProfileContractorGetParams) (api.V1OrganizationsOrganizationIDProfileContractorGetRes, error)
	V1OrganizationsOrganizationIDProfileCustomerGet(ctx context.Context, params api.V1OrganizationsOrganizationIDProfileCustomerGetParams) (api.V1OrganizationsOrganizationIDProfileCustomerGetRes, error)

	V1OrganizationsPortfolioPortfolioIDDelete(ctx context.Context, params api.V1OrganizationsPortfolioPortfolioIDDeleteParams) (api.V1OrganizationsPortfolioPortfolioIDDeleteRes, error)
	V1OrganizationsPortfolioPortfolioIDPut(ctx context.Context, req *api.V1OrganizationsPortfolioPortfolioIDPutReq, params api.V1OrganizationsPortfolioPortfolioIDPutParams) (api.V1OrganizationsPortfolioPortfolioIDPutRes, error)
	V1OrganizationsOrganizationIDPortfolioPost(ctx context.Context, req *api.V1OrganizationsOrganizationIDPortfolioPostReq, params api.V1OrganizationsOrganizationIDPortfolioPostParams) (api.V1OrganizationsOrganizationIDPortfolioPostRes, error)
	V1OrganizationsOrganizationIDPortfolioGet(ctx context.Context, params api.V1OrganizationsOrganizationIDPortfolioGetParams) (api.V1OrganizationsOrganizationIDPortfolioGetRes, error)

	V1OrganizationsFavouritesFavouriteIDDelete(ctx context.Context, params api.V1OrganizationsFavouritesFavouriteIDDeleteParams) (api.V1OrganizationsFavouritesFavouriteIDDeleteRes, error)
	V1OrganizationsOrganizationIDFavouritesGet(ctx context.Context, params api.V1OrganizationsOrganizationIDFavouritesGetParams) (api.V1OrganizationsOrganizationIDFavouritesGetRes, error)
	V1OrganizationsOrganizationIDFavouritesPost(ctx context.Context, req *api.V1OrganizationsOrganizationIDFavouritesPostReq, params api.V1OrganizationsOrganizationIDFavouritesPostParams) (api.V1OrganizationsOrganizationIDFavouritesPostRes, error)
}

type Suggest interface {
	V1SuggestCompanyGet(ctx context.Context, params api.V1SuggestCompanyGetParams) (api.V1SuggestCompanyGetRes, error)
	V1SuggestCityGet(ctx context.Context, params api.V1SuggestCityGetParams) (api.V1SuggestCityGetRes, error)
}

type Verification interface {
	V1VerificationsRequestIDAprovePost(ctx context.Context, params api.V1VerificationsRequestIDAprovePostParams) (api.V1VerificationsRequestIDAprovePostRes, error)
	V1VerificationsRequestIDDenyPost(ctx context.Context, req *api.V1VerificationsRequestIDDenyPostReq, params api.V1VerificationsRequestIDDenyPostParams) (api.V1VerificationsRequestIDDenyPostRes, error)
	V1VerificationsRequestIDGet(ctx context.Context, params api.V1VerificationsRequestIDGetParams) (api.V1VerificationsRequestIDGetRes, error)

	V1VerificationsTendersGet(ctx context.Context, params api.V1VerificationsTendersGetParams) (api.V1VerificationsTendersGetRes, error)

	V1VerificationsAdditionsGet(ctx context.Context, params api.V1VerificationsAdditionsGetParams) (api.V1VerificationsAdditionsGetRes, error)

	V1VerificationsOrganizationsOrganizationIDPost(ctx context.Context, req []api.Attachment, params api.V1VerificationsOrganizationsOrganizationIDPostParams) (api.V1VerificationsOrganizationsOrganizationIDPostRes, error)
	V1VerificationsOrganizationsOrganizationIDGet(ctx context.Context, params api.V1VerificationsOrganizationsOrganizationIDGetParams) (api.V1VerificationsOrganizationsOrganizationIDGetRes, error)
	V1VerificationsOrganizationsGet(ctx context.Context, params api.V1VerificationsOrganizationsGetParams) (api.V1VerificationsOrganizationsGetRes, error)

	V1VerificationsQuestionAnswerGet(ctx context.Context, params api.V1VerificationsQuestionAnswerGetParams) (api.V1VerificationsQuestionAnswerGetRes, error)
}

type Employee interface {
	V1EmployeePost(ctx context.Context, req *api.V1EmployeePostReq) (api.V1EmployeePostRes, error)
}

type Questionnaire interface {
	V1QuestionnaireOrganizationIDPost(ctx context.Context, req *api.V1QuestionnaireOrganizationIDPostReq, params api.V1QuestionnaireOrganizationIDPostParams) (api.V1QuestionnaireOrganizationIDPostRes, error)
	V1QuestionnaireOrganizationIDStatusGet(ctx context.Context, params api.V1QuestionnaireOrganizationIDStatusGetParams) (api.V1QuestionnaireOrganizationIDStatusGetRes, error)
	V1QuestionnaireGet(ctx context.Context, params api.V1QuestionnaireGetParams) (api.V1QuestionnaireGetRes, error)
}

type RouterParams struct {
	Error         Error
	Auth          Auth
	Tenders       Tenders
	Catalog       Catalog
	Users         Users
	Survey        Survey
	Organization  Organization
	Suggest       Suggest
	Verification  Verification
	Employee      Employee
	Questionnaire Questionnaire
}

func NewRouter(params RouterParams) *Router {
	return &Router{
		Auth:          params.Auth,
		Error:         params.Error,
		Tenders:       params.Tenders,
		Catalog:       params.Catalog,
		Users:         params.Users,
		Survey:        params.Survey,
		Organization:  params.Organization,
		Suggest:       params.Suggest,
		Verification:  params.Verification,
		Employee:      params.Employee,
		Questionnaire: params.Questionnaire,
	}
}
