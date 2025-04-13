package service

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type QuestionnaireCreateParams struct {
	OrganizationID int
	Answers        models.Answers
	IsCompleted    bool
}

type QuestionnaireGetParams struct {
	Offset     uint64
	Limit      uint64
}
