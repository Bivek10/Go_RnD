package services

import (
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
	"github.com/bivek/fmt_backend/utils"
	"gorm.io/gorm"
)

type ChoiceServices struct {
	repository repository.ChoiceRepository
}

func NewChoiceServices(repo repository.ChoiceRepository) ChoiceServices {
	return ChoiceServices{
		repository: repo,
	}
}

func (c ChoiceServices) WithTrx(trxHandle *gorm.DB) ChoiceServices {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}
func (c ChoiceServices) CreateChoices(choices models.Choices) error {
	err := c.repository.CreateChoices(choices)
	return err
}

func (c ChoiceServices) GetAllChoices(pagination utils.Pagination) ([]models.Choices, int64, error) {
	return c.repository.GetAllChoices(pagination)
}

func (c ChoiceServices) GetChoicesByQuestionID(pagination utils.Pagination, question_id int64) ([]models.Choices, int64, error) {
	return c.repository.GetChoicesByQuestionID(pagination, question_id)
}
