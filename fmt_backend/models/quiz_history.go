package models

type QuizHistory struct {
	Base
	Quiz_ID int64  `json:"quiz_id"`
	User_ID string `json:"user_id"`
	Score   int64  `json:"score"`
}

func (m QuizHistory) TableName() string {
	return "quizhistory"
}

// to map convert questions to map
func (m QuizHistory) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"quiz_id": m.Quiz_ID,
		"user_id": m.User_ID,
		"score":   m.Score,
	}
}
