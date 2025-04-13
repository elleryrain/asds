package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

type Addition struct {
	VerificationObject

	ID                 int
	TenderID           int
	Title              string
	Content            string
	Attachments        []string
	VerificationStatus VerificationStatus
	CreatedAt          time.Time
}

func ConvertAdditionModelToApi(a Addition) api.Addition {
	return api.Addition{
		ID:                 a.ID,
		TenderID:           a.TenderID,
		Title:              a.Title,
		Content:            a.Content,
		Attachments:        a.Attachments,
		VerificationStatus: string(a.VerificationStatus.ToAPI()),
		CreatedAt:          a.CreatedAt,
	}
}

func (c Addition) ToVerificationObject() api.VerificationRequestObject {
	return api.VerificationRequestObject{
		Type:     api.AdditionVerificationRequestObject,
		Addition: ConvertAdditionModelToApi(c),
	}
}
