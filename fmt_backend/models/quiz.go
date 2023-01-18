package models

type Quizs struct {
	Base
	Quiz_ID int64  `json:"quiz_id"`
	Quiz    string `json:"quiz"`
	
}

func (m Quizs) TableName() string {
	return "quizs"
}

func (m Quizs) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"quiz_id": m.Quiz_ID,
		"quiz":   m.Quiz,
		
	}
}
