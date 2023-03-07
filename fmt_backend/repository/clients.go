package repository

import (
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewClientRepository(db infrastructure.Database, loggger infrastructure.Logger) ClientRepository {
	return ClientRepository{
		db:     db,
		logger: loggger,
	}
}

func (c ClientRepository) WithTrx(trxHandle *gorm.DB) ClientRepository {
	if trxHandle == nil {
		c.logger.Zap.Error("Transcation Database no found in gin context.")
		return c
	}
	c.db.DB = trxHandle
	return c
}

func (c ClientRepository) CreateClient(Clients models.Clients) error {

	err := c.db.DB.Create(&Clients).Error

	return err
}

func (c ClientRepository) LoginClient(Email string) (models.Clients, error) {
	clients := models.Clients{}
	queryBuilder := c.db.DB
	queryBuilder = queryBuilder.Model(&models.Clients{})
	queryBuilder.Where("email = ?", Email)
	err := queryBuilder.Find(&clients).Error
	return clients, err
}
