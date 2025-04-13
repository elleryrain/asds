package mail

import (
	"html/template"
	"strings"

	"git.ubrato.ru/ubrato/dispatch-service/internal/broker"
	"git.ubrato.ru/ubrato/dispatch-service/internal/pb/models/v1"
	"google.golang.org/protobuf/proto"
)

func (s *Service) emailConfirmHandler(msg *broker.Message) error {
	emailConfirmation := &models.PasswordRecovery{}
	err := proto.Unmarshal(msg.Data, emailConfirmation)
	if err != nil {
		return err
	}
	template, err := template.ParseFiles("./templates/emailConfirm.templ")
	if err != nil {
		return err
	}
	body := new(strings.Builder)
	template.Execute(body, emailConfirmation)
	return s.smtpClient.Send(emailConfirmation.Email, "Подтвеждение почты", []byte(body.String()))
}
