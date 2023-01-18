package models

type QuestionChoices struct {
	Base
	Q_ID int64 `json:"q_id"`
	C_ID int64 `json:"c_id"`
}

func (m QuestionChoices) TableName() string {
	return "questionchoices"
}
