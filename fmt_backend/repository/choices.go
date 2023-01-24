package repository

import (
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
	"gorm.io/gorm"
)

type ChoiceRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewChoiceRepository(db infrastructure.Database, logger infrastructure.Logger) ChoiceRepository {
	return ChoiceRepository{
		db:     db,
		logger: logger,
	}
}

func (q ChoiceRepository) WithTrx(trxHandle *gorm.DB) ChoiceRepository {
	if trxHandle == nil {
		q.logger.Zap.Error("Transction Database not found in gin context")
		return q
	}
	q.db.DB = trxHandle
	return q
}

// save -> quiz data.
func (q ChoiceRepository) CreateChoices(Choices models.Choices) error {
	return q.db.DB.Create(&Choices).Error
}

func (c ChoiceRepository) GetAllChoices(pagination utils.Pagination) ([]models.Choices, int64, error) {
	var choices []models.Choices
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.Choices{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`choices`.`name` LIKE ?", searchQuery))
	}
	err := queryBuilder.Find(&choices).Offset(-1).Limit(-1).Count(&totalRows).Error
	return choices, totalRows, err
}

func (c ChoiceRepository) GetChoicesByQuestionID(pagination utils.Pagination, question_id int64) ([]models.Choices, int64, error) {
	var choices []models.Choices
	var totalRows int64 = 0
	var err error
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.Choices{})
	queryBuilder.Where(&models.Choices{Question_ID: question_id})
	err = queryBuilder.Find(&choices).Error
	return choices, totalRows, err
}
