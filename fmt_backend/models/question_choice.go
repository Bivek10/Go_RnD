package models

import "gorm.io/gorm"

// type QuestionChoices struct {
// 	Base
// 	Q_ID int64 `json:"q_id"`
// 	C_ID int64 `json:"c_id"`
// }

// func (m QuestionChoices) TableName() string {
// 	return "questionchoices"
// }

type QuestionChoices struct {
	gorm.Model
	Base
	Q_ID     int64     `json:"q_id"`
	Question string    `json:"question"`
	Quiz_ID  int64     `json:"quiz_id"`
	Choices  []Choices `gorm:"many2many:choices:"`
}
