package dto

type CreateContactsRequest []CreateContactRequest

type CreateContactRequest struct {
	Name               string              `json:"name"`
	FirstName          string              `json:"first_name"`
	LastName           string              `json:"last_name"`
	ResponsibleUserID  int                 `json:"responsible_user_id,omitempty"`
	CustomFieldsValues []CustomFieldsValue `json:"custom_fields_values"`
}

type CreateContactsResponse struct {
	Contacts []CreateContactsResponseContacts `json:"contacts"`
}

type CreateContactsResponseContacts struct {
	ID         int  `json:"id"`
	IsDeleted  bool `json:"is_deleted"`
	IsUnsorted bool `json:"is_unsorted"`
}
