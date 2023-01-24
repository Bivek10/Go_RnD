package services

import (
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/repository"
	"github.com/bivek/fmt_backend/utils"
	"gorm.io/gorm"
)

type QuizHistoryServices struct {
	repository repository.QuizHistoryRepository
}

func NewQuizHistoryServices(repo repository.QuizHistoryRepository) QuizHistoryServices {
	return QuizHistoryServices{
		repository: repo,
	}
}

func (c QuizHistoryServices) WithTrx(trxHandle *gorm.DB) QuizHistoryServices {
	c.repository = c.repository.WithTrx(trxHandle)
	return c
}

func (c QuizHistoryServices) CreateHistory(quizhistory models.QuizHistory) error {
	err := c.repository.CreateQuizHistory(quizhistory)
	return err
}

func (c QuizHistoryServices) GetAllHistory(pagination utils.Pagination) ([]models.QuizHistory, int64, error) {
	return c.repository.GetAllHistory(pagination)
}

func (c QuizHistoryServices) GetHistoryByUserID(pagination utils.Pagination, user_id int) ([]models.QuizHistory, int64, error) {
	return c.repository.GetHistoryByUserID(pagination, int64(user_id))
}
