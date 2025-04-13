package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
)

type Answer struct {
	Answer  interface{}      `json:"answer"`
	Comment Optional[string] `json:"comment"`
}

func ConvertAPIToAnswer(apiAnswer api.QuestionnaireAnswer) Answer {
	answer := Answer{
		Comment: Optional[string]{Value: apiAnswer.Comment.Value, Set: apiAnswer.Comment.Set},
	}

	switch apiAnswer.Answer.Type {
	case api.StringQuestionnaireAnswerAnswer:
		if value, ok := apiAnswer.Answer.GetString(); ok {
			answer.Answer = value
		}

	case api.StringArrayQuestionnaireAnswerAnswer:
		if value, ok := apiAnswer.Answer.GetStringArray(); ok {
			answer.Answer = value
		}
	}

	return answer
}

func ConvertAnswerToAPI(answer Answer) api.QuestionnaireAnswer {
	apiAnswer := api.QuestionnaireAnswer{
		Comment: api.OptString{Value: answer.Comment.Value, Set: answer.Comment.Set},
	}

	switch v := answer.Answer.(type) {
	case []interface{}:
		apiAnswer.Answer = api.NewStringArrayQuestionnaireAnswerAnswer(
			convert.Slice[[]interface{}, []string](v, func(i interface{}) string {
				if s, ok := i.(string); ok {
					return s
				}
				return ""
			}),
		)
	case string:
		apiAnswer.Answer = api.NewStringQuestionnaireAnswerAnswer(v)
	}

	return apiAnswer
}

type Answers []Answer

func (a Answers) Value() (driver.Value, error) {
	if a == nil {
		return []byte("[]"), nil
	}

	return json.Marshal(a)
}

func (a *Answers) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Questionnaire struct {
	ID           int
	Organization Organization
	Answers      Answers
	IsCompleted  bool
	CompletedAt  Optional[time.Time]
	CreatedAt    time.Time
}

func ConvertQuestionnaireToAPI(que Questionnaire) api.Questionnaire {
	return api.Questionnaire{
		ID:           que.ID,
		Organization: ConvertOrganizationModelToApi(que.Organization),
		Answers:      convert.Slice[Answers, []api.QuestionnaireAnswer](que.Answers, ConvertAnswerToAPI),
		IsCompleted:  que.IsCompleted,
		CompletedAt:  api.OptDateTime{Value: que.CompletedAt.Value, Set: que.CompletedAt.Set},
		CreatedAt:    que.CreatedAt,
	}
}
