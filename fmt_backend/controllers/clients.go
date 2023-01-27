package controllers

import (
	//"fmt"
	"net/http"

	"github.com/bivek/fmt_backend/constants"
	"github.com/bivek/fmt_backend/errors"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/responses"
	"github.com/bivek/fmt_backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ClientController -> struct
type ClientController struct {
	logger          infrastructure.Logger
	clientService   services.ClientService
	firebaseService services.FirebaseService
	env             infrastructure.Env
}

// NewClientController -> constructor
func NewClientController(
	logger infrastructure.Logger,
	clientService services.ClientService,
	firebaseSerivce services.FirebaseService,
	env infrastructure.Env,
) ClientController {
	return ClientController{
		logger:          logger,
		clientService:   clientService,
		firebaseService: firebaseSerivce,
		env:             env,
	}
}

// CreateUser -> Create User
func (cc ClientController) CreateClients(c *gin.Context) {
	clients := models.Clients{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB) // explicitly define the value type..

	if err := c.ShouldBindJSON(&clients); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind user data")
		responses.HandleError(c, err)
		return
	}

	clientResposne, message, err := cc.clientService.WithTrx(trx).CreateClient(clients)
	if message != "" {
		responses.SuccessJSON(c, http.StatusOK, message)
	}
	if clientResposne.AccessToken == "" {
		responses.SuccessJSON(c, http.StatusOK, "Failed to generate access_token")
	}
	if clientResposne.RefreshToken == "" {
		responses.SuccessJSON(c, http.StatusOK, "Failed to generate refresht_token")
	}

	if err != nil {
		cc.logger.Zap.Error("Error [CreateUser] [db CreateUser]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create err ")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "User Created Sucessfully")

}
