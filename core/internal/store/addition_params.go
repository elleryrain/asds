package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type AdditionCreateParams struct {
	TenderID    int
	Title       string
	Content     string
	Attachments []string
}

type AdditionGetParams struct {
	TenderID     models.Optional[int]
	AdditionIDs  models.Optional[[]int]
	VerifiedOnly bool
}

type AdditionUpdateVerifStatusParams struct {
	AdditionID         int
	VerificationStatus models.VerificationStatus
}
