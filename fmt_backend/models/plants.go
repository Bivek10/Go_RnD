package models

type Plant struct {
	Base
	PlantId        string `json:"plant_id"`
	PlantName      string `json:"plant_name"`
	ScientificName string `json:"plant_scientific_name"`
	PlantType      string `json:"plant_type"`
	Temperature    string `json:"temperature"`
	Description    string `json:"description"`
	ImageUrl       string `json:"image_url"`
	GrowingSession string `json:"growing_season"`
}

func (m Plant) TableName() string {
	return "plants"
}

// to map convert plants to map
func (m Plant) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"plant_id":         m.PlantId,
		"plant_name":       m.PlantName,
		"scientific_name":  m.ScientificName,
		"plant_type":       m.PlantType,
		"temperature":      m.Temperature,
		"description":      m.Description,
		"image_url":        m.ImageUrl,
		"growing_season": m.GrowingSession,
	}
}
