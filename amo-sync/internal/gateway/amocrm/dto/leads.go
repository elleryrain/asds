package dto

type CreateLeadsRequest []CreateLeadRequest

type CreateLeadRequest struct {
	Name               string                       `json:"name"`
	StatusID           int                          `json:"status_id"`
	PipelineID         int                          `json:"pipeline_id"`
	ResponsibleUserID  int                          `json:"responsible_user_id,omitempty"`
	CustomFieldsValues []CustomFieldsValue          `json:"custom_fields_values,omitempty"`
	TagsToAdd          []CreateLeadRequestTagsToAdd `json:"tags_to_add,omitempty"`
	Embedded           CreateLeadRequestEmbedded    `json:"_embedded"`
}

type CreateLeadRequestEmbedded struct {
	Contacts  []CreateLeadRequestEmbeddedContacts  `json:"contacts"`
	Companies []CreateLeadRequestEmbeddedCompanies `json:"companies"`
}

type CreateLeadRequestTagsToAdd struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateLeadRequestEmbeddedContacts struct {
	ID int `json:"id"`
}

type CreateLeadRequestEmbeddedCompanies struct {
	ID int `json:"id"`
}

type CreateLeadsResponse struct {
	Leads []CreateLeadsResponseLeads `json:"leads"`
}

type CreateLeadsResponseLeads struct {
	ID int `json:"id"`
}
