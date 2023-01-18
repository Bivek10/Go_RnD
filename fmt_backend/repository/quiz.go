package repository

import (
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
	"gorm.io/gorm"
)

type QuizRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewQuizRepository(db infrastructure.Database, logger infrastructure.Logger) QuizRepository {
	return QuizRepository{
		db:     db,
		logger: logger,
	}
}

func (q QuizRepository) WithTrx(trxHandle *gorm.DB) QuizRepository {
	if trxHandle == nil {
		q.logger.Zap.Error("Transction Database not found in gin context")
		return q
	}
	q.db.DB = trxHandle
	return q
}

// save -> quiz data.
func (q QuizRepository) CreateQuiz(Quizs models.Quizs) error {
	return q.db.DB.Create(&Quizs).Error
}

func (c QuizRepository) GetAllQuiz(pagination utils.Pagination) ([]models.Quizs, int64, error) {
	var quizs []models.Quizs
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.Quizs{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`quizs`.`name` LIKE ?", searchQuery))
	}
	err := queryBuilder.Find(&quizs).Offset(-1).Limit(-1).Count(&totalRows).Error
	return quizs, totalRows, err
}
