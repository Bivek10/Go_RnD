package models

type Choices struct {
	Base
	C_ID      int64  `json:"c_id"`
	Choice    string `json:"choice"`
	IsCorrect bool   `json:"isCorrect"`
	Q_ID      int64  `json:"q_id"`
}

func (m Choices) TableName() string {
	return "choices"
}

// to map convert questions to map
func (m Choices) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"c_id":      m.C_ID,
		"choice":    m.Choice,
		"isCorrect": m.IsCorrect,
		"q_id":      m.Q_ID,
	}
}