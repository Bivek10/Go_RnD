package models

type Quizs struct {
	Base
	Quiz string `json:"quiz"`
}

func (m Quizs) TableName() string {
	return "quizs"
}

func (m Quizs) ToMap() map[string]interface{} {
	return map[string]interface{}{

		"quiz": m.Quiz,
	}
}
