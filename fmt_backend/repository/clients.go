package repository

import (
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/utils"
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

func (c ClientRepository) CreateClient(Clients models.Clients) (models.ClientRequestResponse, string, error) {
	clientrequestresponse := models.ClientRequestResponse{}
	clients := models.Clients{}
	print(clientrequestresponse)
	queryBuilder := c.db.DB
	queryBuilder = queryBuilder.Model(&models.Clients{}).Where(&models.User{Email: Clients.Email}).Find(clients)
	var message string = ""
	var err error

	if clients.Email != "" {
		message = "Email already exist!"
	} else {
		clients.Password = utils.EncryptPassword([]byte(clients.Password))
		err = c.db.DB.Create(clients).Error
		clientrequestresponse.FirstName = clients.FirstName
		clientrequestresponse.LastName = clients.LastName
		clientrequestresponse.Address = clients.Address
		clientrequestresponse.Email = clients.Email
		accesstoken, err, refreshtoken, refresherror := utils.GenerateJWT(clients.Email)
		if err != nil {
			return clientrequestresponse, message, err
		}
		if refresherror != nil {
			return clientrequestresponse, message, err
		}
		clientrequestresponse.AccessToken = accesstoken
		clientrequestresponse.RefreshToken = refreshtoken

	}

	return clientrequestresponse, message, err
}
