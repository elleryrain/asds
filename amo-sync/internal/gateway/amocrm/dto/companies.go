package dto

type CreateCompaniesRequest []CreateCompanyRequest

type CreateCompanyRequest struct {
	Name               string              `json:"name"`
	ResponsibleUserID  int                 `json:"responsible_user_id,omitempty"`
	CustomFieldsValues []CustomFieldsValue `json:"custom_fields_values"`
}

type CreateCompaniesResponse struct {
	Companies []CreateCompaniesResponseCompanies `json:"companies"`
}

type CreateCompaniesResponseCompanies struct {
	ID         int  `json:"id"`
	IsDeleted  bool `json:"is_deleted"`
	IsUnsorted bool `json:"is_unsorted"`
}
