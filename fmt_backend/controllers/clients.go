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
	"github.com/bivek/fmt_backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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
	println("hello new routes")
	clients := models.Clients{}
	clientrequestresponse := models.ClientRequestResponse{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB) // explicitly define the value type..

	if err := c.ShouldBindJSON(&clients); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind user data")
		responses.HandleError(c, err)
		return
	}
	// encrypt password.
	clients.Password = utils.EncryptPassword([]byte(clients.Password))

	if err := cc.clientService.WithTrx(trx).CreateClient(clients); err != nil {

		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			if mysqlError.Number == 1062 {
				err := errors.Conflict.Wrap(err, "Email already exits!")
				errs := errors.SetCustomMessage(err, "Email already exists!")
				responses.HandleError(c, errs)
				return
			}
		}
		cc.logger.Zap.Error("Error [CreateClient user] [db Clientuser]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create client user")
		responses.HandleError(c, err)
		return
	}

	clientrequestresponse.FirstName = clients.FirstName
	clientrequestresponse.LastName = clients.LastName
	clientrequestresponse.Address = clients.Address
	clientrequestresponse.Email = clients.Email

	accesstoken, err, refreshtoken, refresherror := utils.GenerateJWT(clients.Email)
	if err != nil {
		cc.logger.Zap.Error("Error creating the accesstoken: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create access token")
		responses.HandleError(c, err)
		return
	}
	
	if refresherror != nil {
		cc.logger.Zap.Error("Error creating the refreshtoken: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create refresht token")
		responses.HandleError(c, err)
		return
	}
	clientrequestresponse.AccessToken = accesstoken
	clientrequestresponse.RefreshToken = refreshtoken

	responses.SuccessJSON(c, http.StatusOK, clientrequestresponse)

}
