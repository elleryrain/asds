package models

import (
	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

type Winners struct {
	ID           int
	TenderID     int
	Accepted     AcceptedStatus
	Organization Organization
}

func ConvertWinnerModelToApi(winners Winners) api.Winners {
	return api.Winners{
		ID:           winners.ID,
		TenderID:     winners.TenderID,
		Accepted:     winners.Accepted.ToAPI(),
		Organization: ConvertOrganizationModelToApi(winners.Organization),
	}
}

type AcceptedStatus int

const (
	AcceptedStatusUnverified AcceptedStatus = iota
	AcceptedStatusAwaiting
	AcceptedStatusDeclined
	AcceptedStatusApproved
)

var mapAcceptedStatus = map[AcceptedStatus]api.Accepted{
	AcceptedStatusUnverified: api.AcceptedUnverified,
	AcceptedStatusAwaiting:   api.AcceptedAwaiting,
	AcceptedStatusDeclined:   api.AcceptedDeclined,
	AcceptedStatusApproved:   api.AcceptedApproved,
}

func (v AcceptedStatus) ToAPI() api.Accepted {
	return mapAcceptedStatus[v]
}

var mapApiToAcceptedStatus = map[api.Accepted]AcceptedStatus{
	api.AcceptedUnverified: AcceptedStatusUnverified,
	api.AcceptedAwaiting:   AcceptedStatusAwaiting,
	api.AcceptedDeclined:   AcceptedStatusDeclined,
	api.AcceptedApproved:   AcceptedStatusApproved,
}

func APIToAcceptedStatus(apiStatus api.Accepted) AcceptedStatus {
	status, ok := mapApiToAcceptedStatus[apiStatus]
	if !ok {
		return AcceptedStatusUnverified
	}
	return status
}
