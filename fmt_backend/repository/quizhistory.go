package repository

import (
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
	"gorm.io/gorm"
)

type QuizHistoryRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewQuizHistoryRepository(db infrastructure.Database, logger infrastructure.Logger) QuizHistoryRepository {
	return QuizHistoryRepository{
		db:     db,
		logger: logger,
	}
}

func (q QuizHistoryRepository) WithTrx(trxHandle *gorm.DB) QuizHistoryRepository {
	if trxHandle == nil {
		q.logger.Zap.Error("Transction Database not found in gin context")
		return q
	}
	q.db.DB = trxHandle
	return q
}

// save -> quiz data.
func (q QuizHistoryRepository) CreateQuizHistory(QuizHistory models.QuizHistory) error {
	return q.db.DB.Create(&QuizHistory).Error
}

func (c QuizHistoryRepository) GetAllHistory(pagination utils.Pagination) ([]models.QuizHistory, int64, error) {
	var quizhistory []models.QuizHistory
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.Choices{})

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`quizhistory`.`name` LIKE ?", searchQuery))
	}
	err := queryBuilder.Find(&quizhistory).Offset(-1).Limit(-1).Count(&totalRows).Error
	return quizhistory, totalRows, err
}

func (c QuizHistoryRepository) GetHistoryByUserID(pagination utils.Pagination, user_id int64) ([]models.QuizHistory, int64, error) {
	var quizhistory []models.QuizHistory
	var totalRows int64 = 0
	var err error
	queryBuilder := c.db.DB.Limit(pagination.PageSize).Offset(pagination.Offset).Order("created_at desc")
	queryBuilder = queryBuilder.Model(&models.QuizHistory{})
	queryBuilder.Where(&models.QuizHistory{User_ID: user_id})
	err = queryBuilder.Find(&quizhistory).Error
	return quizhistory, totalRows, err
}
