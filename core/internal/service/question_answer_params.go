package service

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type CreateQuestionAnswerParams struct {
	TenderID             int
	AuthorOrganizationID int
	ParentID             models.Optional[int]
	Type                 models.QuestionAnswerType
	Content              string
}