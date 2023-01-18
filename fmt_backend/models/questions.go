package models

type Questions struct {
	Base
	Q_ID     int64  `json:"q_id"`
	Question string `json:"question"`
	Quiz_ID  int64  `json:"quiz_id"`
}

func (m Questions) TableName() string {
	return "questions"
}

// to map convert questions to map
func (m Questions) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"q_id":     m.Q_ID,
		"question": m.Question,
		"quiz_id":  m.Quiz_ID,
	}
}
