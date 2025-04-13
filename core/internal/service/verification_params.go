package service

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type VerificationRequestCreateParams struct {
	ObjectID    int
	ObjectType  models.ObjectType
	Attachments models.Attachments
}

type VerificationRequestUpdateStatusParams struct {
	UserID        int
	RequesID      int
	Status        models.VerificationStatus
	ReviewComment models.Optional[string]
}

type VerificationObjectUpdateStatusParams struct {
	Object models.VerificationObject
	Status models.VerificationStatus
}

type VerificationRequestsObjectGetParams struct {
	ObjectType models.ObjectType
	ObjectID   models.Optional[int]
	Status     []models.VerificationStatus
	Page       uint64
	PerPage    uint64
}
