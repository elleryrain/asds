package models

import api "gitlab.ubrato.ru/ubrato/notification/api/gen"

var (
	StatusBlockInReview = &StatusBlock{
		Text:   "На модерации",
		Status: StatusInReview,
	}

	StatusBlockDeclined = &StatusBlock{
		Text:   "Верификация не пройдена",
		Status: StatusDeclined,
	}

	StatusBlockOrganizationApproved = &StatusBlock{
		Text:   "Исправьте и отправьте заново",
		Status: StatusDeclined,
	}

	StatusBlockApproved = &StatusBlock{
		Text:   "Верификация пройдена",
		Status: StatusApproved,
	}
)

type StatusBlock struct {
	Text   string `json:"text,omitempty"`
	Status Status `json:"status,omitempty"`
}

func NewOptStatusBlock(sb *StatusBlock) Optional[StatusBlock] {
	if sb == nil {
		return Optional[StatusBlock]{}
	}

	return NewOptional(*sb)
}

func ConvertStatusBlockToOptAPI(statusBlock *StatusBlock) api.OptStatusBlock {
	if statusBlock == nil {
		return api.OptStatusBlock{}
	}

	return api.NewOptStatusBlock(api.StatusBlock{
		Text:   statusBlock.Text,
		Status: statusBlock.Status.ToAPI(),
	})
}

type Status int

const (
	StatusInvalid Status = iota
	StatusInReview
	StatusDeclined
	StatusApproved
)

var mapStatusToAPI = map[Status]api.Status{
	StatusInvalid:  api.StatusInvalid,
	StatusInReview: api.StatusInReview,
	StatusDeclined: api.StatusDeclined,
	StatusApproved: api.StatusApproved,
}

func (s Status) ToAPI() api.Status {
	return mapStatusToAPI[s]
}

var mapAPIToStatus = map[api.Status]Status{
	api.StatusInvalid:  StatusInvalid,
	api.StatusInReview: StatusInReview,
	api.StatusDeclined: StatusDeclined,
	api.StatusApproved: StatusApproved,
}

func APIToStatus(apiStatus api.Status) Status {
	status, ok := mapAPIToStatus[apiStatus]
	if !ok {
		return StatusInvalid
	}
	return status
}
