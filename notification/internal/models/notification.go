package models

import (
	"fmt"

	api "gitlab.ubrato.ru/ubrato/notification/api/gen"
	modelsv1 "gitlab.ubrato.ru/ubrato/notification/internal/models/gen/proto/models/v1"
)

type Notification struct {
	ID           int           `json:"id"`
	UserID       int           `json:"user_id"`
	Title        string        `json:"title"`
	Comment      string        `json:"comment,omitempty"`
	ActionButton *ActionButton `json:"action_button,omitempty"`
	StatusBlock  *StatusBlock  `json:"status_block,omitempty"`
	IsRead       bool          `json:"is_read"`
}

func ConvertNotificationToAPI(notification Notification) api.Notification {
	return api.Notification{
		ID:           notification.ID,
		UserID:       notification.UserID,
		Title:        notification.Title,
		Comment:      api.OptString{Value: notification.Comment, Set: notification.Comment != ""},
		ActionButton: ConvertActionButtonToOptAPI(notification.ActionButton),
		StatusBlock:  ConvertStatusBlockToOptAPI(notification.StatusBlock),
		IsRead:       notification.IsRead,
	}
}

// Верификация организации
func MustOrganizationVerificationNotify(eventNotify *modelsv1.Notification) Notification {
	notify := Notification{
		UserID: int(eventNotify.User.Id),
		Title:  "Верификация организации",
	}

	switch eventNotify.Verification.Status {
	case modelsv1.Status_STATUS_IN_REVIEW:
		notify.Comment = "Документы отправлены на модерацию. Пожалуйста, ожидайте."
		notify.StatusBlock = StatusBlockInReview

	case modelsv1.Status_STATUS_DECLINED:
		notify.Comment = fmt.Sprintf("Верификация не пройдена.\n%s", eventNotify.Verification.Comment)
		notify.StatusBlock = StatusBlockDeclined
		notify.ActionButton = &ActionButton{
			Text: "Перейти",
			Url:  "/profile/documents",
		}

	case modelsv1.Status_STATUS_APPROVED:
		if eventNotify.User.IsContractor {
			notify.Comment = "Верификация пройдена, теперь можно откликаться на тендеры."
		} else {
			notify.Comment = "Верификация пройдена, теперь можно создавать тендеры."
		}

		notify.StatusBlock = StatusBlockApproved
	}

	return notify
}

// Верификация тендеров
func MustTenderVerificationNotify(eventNotify *modelsv1.Notification) Notification {
	notify := Notification{
		Title:  "Создание тендера",
		UserID: int(eventNotify.User.Id),
		ActionButton: &ActionButton{
			Text: "Перейти",
			Url:  fmt.Sprintf("/tender/%v", eventNotify.Object.Id),
		},
	}

	switch eventNotify.Verification.Status {
	case modelsv1.Status_STATUS_IN_REVIEW:
		notify.Comment = fmt.Sprintf("Заявка на создание тендера № %v \"%v\" отправлена на модерацию. Пожалуйста, ожидайте.", eventNotify.Object.Id, eventNotify.Object.Tender.Title)
		notify.StatusBlock = StatusBlockInReview

	case modelsv1.Status_STATUS_DECLINED:
		notify.Comment = fmt.Sprintf("Тендер № %v \"%v\" не прошел модерацию. Внесите исправления.\n%s", eventNotify.Object.Id, eventNotify.Object.Tender.Title, eventNotify.Verification.Comment)
		notify.StatusBlock = StatusBlockDeclined

	case modelsv1.Status_STATUS_APPROVED:
		startTime := eventNotify.Object.Tender.ReceptionStart.AsTime()
		notify.Comment = fmt.Sprintf(
			"Тендер № %v \"%v\" прошел модерацию, прием откликов начнется %s в %s.",
			eventNotify.Object.Id, eventNotify.Object.Tender.Title,
			startTime.Format("02/01/2006"), startTime.Format("15:04"))

		notify.StatusBlock = StatusBlockApproved
	}

	return notify
}

// Верификация доп информации для тендера
func MustAdditionVerificationNotify(eventNotify *modelsv1.Notification) Notification {
	notify := Notification{
		UserID: int(eventNotify.User.Id),
		Title:  "Верификация доп. информации",
		ActionButton: &ActionButton{
			Text: "Перейти",
			Url:  fmt.Sprintf("/tender/%v/more_inforamtion", eventNotify.Object.Tender.Id),
		},
	}

	switch eventNotify.Verification.Status {
	case modelsv1.Status_STATUS_IN_REVIEW:
		notify.Comment = fmt.Sprintf("Доп. информация для тендера № %v \"%v\" отправлена на модерацию. Пожалуйста, ожидайте.", eventNotify.Object.Tender.Id, eventNotify.Object.Tender.Title)
		notify.StatusBlock = StatusBlockInReview

	case modelsv1.Status_STATUS_DECLINED:
		notify.Comment = fmt.Sprintf("Верификация доп. информации для тендера № %v \"%v\" не пройдена.\n%s", eventNotify.Object.Tender.Id, eventNotify.Object.Tender.Title, eventNotify.Verification.Comment)
		notify.StatusBlock = StatusBlockDeclined

	case modelsv1.Status_STATUS_APPROVED:
		notify.Comment = fmt.Sprintf("Верификация пройдена, доп. информация для тендера № %v \"%v\" опубликована.", eventNotify.Object.Tender.Id, eventNotify.Object.Tender.Title)
		notify.StatusBlock = StatusBlockApproved
	}

	return notify
}

// Верификация доп информации для тендера
func MustQuestionAnswerVerificationNotify(eventNotify *modelsv1.Notification) Notification {
	notify := Notification{
		UserID: int(eventNotify.User.Id),
		Title:  "Верификация вопроса/ответа",
		ActionButton: &ActionButton{
			Text: "Перейти",
			Url:  fmt.Sprintf("/tender/%v/questions_and_answers", eventNotify.Object.Tender.Id),
		},
	}

	var qaType string
	if eventNotify.User.IsContractor {
		qaType = "вопрос"
	} else {
		qaType = "ответ"
	}

	switch eventNotify.Verification.Status {
	case modelsv1.Status_STATUS_IN_REVIEW:
		notify.Comment = fmt.Sprintf("Ваш %s для тендера № %v \"%v\" отправлен на модерацию. Пожалуйста, ожидайте.", qaType, eventNotify.Object.Tender.Id, eventNotify.Object.Tender.Title)
		notify.StatusBlock = StatusBlockInReview

	case modelsv1.Status_STATUS_DECLINED:
		notify.Comment = fmt.Sprintf("Верификация %sа для тендера № %v \"%v\" не пройдена.\n%s", qaType, eventNotify.Object.Tender.Id, eventNotify.Object.Tender.Title, eventNotify.Verification.Comment)
		notify.StatusBlock = StatusBlockDeclined

	case modelsv1.Status_STATUS_APPROVED:
		notify.Comment = fmt.Sprintf("Верификация пройдена, %s для тендера № %v \"%v\" опубликован.", qaType, eventNotify.Object.Tender.Id, eventNotify.Object.Tender.Title)
		notify.StatusBlock = StatusBlockApproved
	}

	return notify
}

func MustQuestionAnswerNotify(eventNotify *modelsv1.Notification) Notification {
	notify := Notification{
		UserID: int(eventNotify.User.Id),
		Title:  "Вопрос/ответ",
		ActionButton: &ActionButton{
			Text: "Перейти",
			Url:  fmt.Sprintf("/tender/%v/questions_and_answers", eventNotify.Object.Tender.Id),
		},
	}

	if eventNotify.User.IsContractor {
		notify.Comment = fmt.Sprintf("На тендер № %v \"%v\" добавлен ответ на ваш вопрос.", eventNotify.Object.Tender.Id, eventNotify.Object.Tender.Title)
	} else {
		notify.Comment = fmt.Sprintf("На тендер № %v \"%v\" поступил запрос на уточнение информации.", eventNotify.Object.Tender.Id, eventNotify.Object.Tender.Title)
	}

	return notify
}

func MustWinnerNotify(eventNotify *modelsv1.Notification) Notification {
	notify := Notification{
		UserID: int(eventNotify.User.Id),
		Title:  "Выбор победителя",
		ActionButton: &ActionButton{
			Text: "Перейти",
			Url:  fmt.Sprintf("/tender/%v/responses", eventNotify.Object.Tender.Id),
		},
	}

	if eventNotify.User.IsContractor {
		notify.Comment = fmt.Sprintf("Ваша компания выбрана победителем тендера № %v \"%v\", не забудьте оставить отзыв", eventNotify.Object.Tender.Id, eventNotify.Object.Tender.Title)
	} else {
		notify.Comment = fmt.Sprintf("Вы выбрали победителя(ей) для тендера № %v \"%v\", не забудьте оставить отзыв(ы)", eventNotify.Object.Tender.Id, eventNotify.Object.Tender.Title)
	}

	return notify
}
