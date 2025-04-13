package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
)

type VerificationRequestPagination[T VerificationObject] struct {
	VerificationRequests []VerificationRequest[T]
	Pagination           pagination.Pagination
}

type Attachment struct {
	Name Optional[string] `json:"name"`
	Url  string           `json:"url"`
}

type Attachments []Attachment

func (a Attachments) Value() (driver.Value, error) {
	if a == nil {
		return []byte("[]"), nil
	}

	return json.Marshal(a)
}

func (a *Attachments) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func ConvertAPIToAttachment(attachment api.Attachment) Attachment {
	return Attachment{
		Name: Optional[string]{Value: attachment.Name.Value, Set: attachment.Name.Set},
		Url:  attachment.URL.String(),
	}
}

func ConvertAttachmentToApi(attachment Attachment) api.Attachment {
	return api.Attachment{
		Name: api.OptString{Value: attachment.Name.Value, Set: attachment.Name.Set},
		URL:  stringToUrl(attachment.Url),
	}
}

type VerificationObject interface {
	ToVerificationObject() api.VerificationRequestObject
}

type VerificationRequest[T VerificationObject] struct {
	ID            int
	Reviewer      EmployeeUser
	ObjectType    ObjectType
	ObjectID      int
	Object        T
	Content       string
	Attachments   Attachments
	Status        VerificationStatus
	ReviewComment string
	CreatedAt     time.Time
	ReviewedAt    time.Time
}

func VerificationRequestModelToApi[T VerificationObject](request VerificationRequest[T]) api.VerificationRequest {
	return api.VerificationRequest{
		ID:            request.ID,
		Reviewer:      api.OptEmployeeUser{Value: ConvertEmployeeUserModelToApi(request.Reviewer), Set: request.Reviewer.ID != 0},
		ObjectType:    api.ObjectType(request.ObjectType.ToAPI()),
		Object:        request.Object.ToVerificationObject(),
		Content:       request.Content,
		Attachments:   convert.Slice[Attachments, []api.Attachment](request.Attachments, ConvertAttachmentToApi),
		Status:        request.Status.ToAPI(),
		ReviewComment: api.OptString{Value: request.ReviewComment, Set: request.ReviewComment != ""},
		CreatedAt:     request.CreatedAt,
		ReviewedAt:    api.OptDateTime{Value: request.ReviewedAt, Set: !request.ReviewedAt.IsZero()},
	}
}

type VerificationStatus int

const (
	VerificationStatusUnverified VerificationStatus = iota
	VerificationStatusInReview
	VerificationStatusDeclined
	VerificationStatusApproved
)

var mapVerificationStatus = map[VerificationStatus]api.VerificationStatus{
	VerificationStatusUnverified: api.VerificationStatusUnverified,
	VerificationStatusInReview:   api.VerificationStatusInReview,
	VerificationStatusDeclined:   api.VerificationStatusDeclined,
	VerificationStatusApproved:   api.VerificationStatusApproved,
}

func (v VerificationStatus) ToAPI() api.VerificationStatus {
	return mapVerificationStatus[v]
}

var mapApiToVerificationStatus = map[api.VerificationStatus]VerificationStatus{
	api.VerificationStatusUnverified: VerificationStatusUnverified,
	api.VerificationStatusInReview:   VerificationStatusInReview,
	api.VerificationStatusDeclined:   VerificationStatusDeclined,
	api.VerificationStatusApproved:   VerificationStatusApproved,
}

func APIToVerificationStatus(apiStatus api.VerificationStatus) VerificationStatus {
	status, ok := mapApiToVerificationStatus[apiStatus]
	if !ok {
		return VerificationStatusUnverified
	}
	return status
}
