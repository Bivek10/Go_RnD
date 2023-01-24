package models

type Questions struct {
	//gorm.Model
	Base
	Question string    `json:"question"`
	Quiz_ID  int64     `json:"quiz_id"`
	Choices  []Choices `json:"choices" gorm:"foreignKey:question_id"`
}

func (m Questions) TableName() string {
	return "questions"
}

// to map convert questions to map
func (m Questions) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"question": m.Question,
		"quiz_id":  m.Quiz_ID,
	}
}
