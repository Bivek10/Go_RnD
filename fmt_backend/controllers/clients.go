package controllers

import (
	//"fmt"
	"net/http"

	"github.com/bivek/fmt_backend/constants"
	"github.com/bivek/fmt_backend/errors"
	"github.com/bivek/fmt_backend/helpers"
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
	jwtServices     services.JWTService
	firebaseService services.FirebaseService
	env             infrastructure.Env
}

// NewClientController -> constructor
func NewClientController(
	logger infrastructure.Logger,
	clientService services.ClientService,
	jwtService services.JWTService,
	firebaseSerivce services.FirebaseService,
	env infrastructure.Env,
) ClientController {
	return ClientController{
		logger:          logger,
		clientService:   clientService,
		firebaseService: firebaseSerivce,
		env:             env,
		jwtServices:     jwtService,
	}
}

// CreateUser -> Create User
func (cc ClientController) CreateClients(c *gin.Context) {
	cc.logger.Zap.Info("client running ============")
	clients := models.BaseClient{}
	clientrequestresponse := models.ClientRequestResponse{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB) // explicitly define the value type..

	imageFile, imageFileHeader, err := c.Request.FormFile("profile_photo")

	if err != nil {
		msg := "Error in getting image from form-file"
		cc.logger.Zap.Error(msg, err)
		err = errors.BadRequest.Wrap(err, msg)
		responses.HandleError(c, err)
		return
	}
	if imageFile == nil || imageFileHeader == nil {
		msg := "Invalid image "
		cc.logger.Zap.Error(msg, err)
		err = errors.BadRequest.Wrap(err, msg)
		responses.HandleError(c, err)
		return
	}

	filename, err := helpers.FileUpload(c, imageFileHeader, constants.ClientImage)

	if err != nil {
		cc.logger.Zap.Error("Error failed to upload photo : ", err)
		err := errors.BadRequest.Wrap(err, "Failed Upload photo")
		responses.HandleError(c, err)
		return
	}

	if err := c.ShouldBind(&clients); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind user data")
		responses.HandleError(c, err)
		return
	}

	clientData := models.Clients{
		BaseClient:   clients,
		ProfilePhoto: filename,
	}

	// encrypt password.
	clients.Password = utils.EncryptPassword([]byte(clients.Password))

	if err := cc.clientService.WithTrx(trx).CreateClient(clientData); err != nil {

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

	accesstoken, err, refreshtoken, refresherror := cc.jwtServices.GenerateJWT(clients.Email)
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

func (cc ClientController) LoginClient(c *gin.Context) {
	authentication := models.Authentication{}
	clientrequestresponse := models.ClientRequestResponse{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB) // explicitly define the value type..

	if err := c.ShouldBindJSON(&authentication); err != nil {
		cc.logger.Zap.Error("Error [login client] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind client authentication data")
		responses.HandleError(c, err)
		return
	}

	clients, err := cc.clientService.WithTrx(trx).LoginClient(authentication.Email)
	if err != nil {
		err := errors.Conflict.Wrap(err, "Email not found")
		errs := errors.SetCustomMessage(err, "Email not found")
		responses.HandleError(c, errs)
		return
	}

	if clients.Email == "" {
		var err error
		err = errors.Conflict.New("Email not found")
		errs := errors.SetCustomMessage(err, "Email not found")
		responses.HandleError(c, errs)
		return
	}
	isMatched := utils.DecryptPassword([]byte(clients.Password), []byte(authentication.Password))
	if isMatched == false {
		cc.logger.Zap.Error("Invalid password")
		err := errors.Conflict.New("Invalid password")
		errs := errors.SetCustomMessage(err, "Invalid password")
		responses.HandleError(c, errs)
		return
	}

	clientrequestresponse.FirstName = clients.FirstName
	clientrequestresponse.LastName = clients.LastName
	clientrequestresponse.Address = clients.Address
	clientrequestresponse.Email = clients.Email
	accesstoken, err, refreshtoken, refresherror := cc.jwtServices.GenerateJWT(clients.Email)

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

func (cc ClientController) ReGenerateClientToken(c *gin.Context) {
	refreshToken := models.RefreshTokenRequest{}
	refreshTokenRequestResponse := models.RefreshTokenRequestResponse{}
	//trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		cc.logger.Zap.Error("Error [Binding token] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind refresh token")
		responses.HandleError(c, err)
		return
	}

	isVerify, err := cc.jwtServices.VerifyRefreshToken(refreshToken.RefreshToken, c)

	if isVerify {
		email := c.MustGet("email")

		access_token, err, refresh_token, refresherror := cc.jwtServices.GenerateJWT(email.(string))

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
		refreshTokenRequestResponse.AcceessToken = access_token
		refreshTokenRequestResponse.RefreshToken = refresh_token
		responses.SuccessJSON(c, http.StatusOK, refreshTokenRequestResponse)

	} else {
		err := errors.Conflict.Wrap(err, "Unauthorized")
		errs := errors.SetCustomMessage(err, "Unauthorized")
		responses.HandleError(c, errs)
		return
	}

}
