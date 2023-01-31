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
	"github.com/golang-jwt/jwt"
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
	accesstoken, err, refreshtoken, refresherror := utils.GenerateJWTV(clients.Email)

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
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		cc.logger.Zap.Error("Error [Binding token] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind refresh token")
		responses.HandleError(c, err)
		return
	}

	//refreshTokens := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJpdmVra2Fya2kxMzNAZ21haWwuY29tIiwiZXhwIjoxNjc1MjQ0NTg2.07VeXSMCda9OJ0WepajGxQvf1_1grh_tnxm8AdknKAk"
	//print(refreshTokens)

	// token, err := jwt.Parse(refreshToken.RefreshToken, func(t *jwt.Token) (interface{}, error) {
	// 	return ([]byte(cc.env.JWRTSecretKey)), nil
	// })
	// clams := jwt.MapClaims{}
	// clams["email"] =
	// clams["exp"] = time.Now().Add(time.Minute * 15).Unix()

	//token, err := jwt.ParseWithClaims(refreshToken.RefreshToken)
	println(refreshToken.RefreshToken)
	token, err := jwt.ParseWithClaims(refreshToken.RefreshToken, &models.SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	println(ok)
		// 	println(token.Header["alg"])
		// 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// }
		return []byte(cc.env.JWRTSecretKey), nil
	})
	// token, err := jwt.Parse(
	// 	refreshToken.RefreshToken, func(token *jwt.Token) (interface{}, error) {
	// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 	println(ok)
	// 	println(token.Header["alg"])
	// 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// }

	// 		println(token.Valid)
	// 		println(cc.env.JWRTSecretKey)
	// 		return []byte(cc.env.JWRTSecretKey), nil
	// 	})
	if err != nil {
		if validationErr, ok := err.(*jwt.ValidationError); ok {
			// Token is malformed
			if validationErr.Errors&jwt.ValidationErrorMalformed != 0 {
				println("Token is malformed")
			} else if validationErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				println("Token is either expired or not active yet")
			} else {
				println(err.Error())
				println("Other err")
			}
		} else {
			println("Other ervvr")
		}
		cc.logger.Zap.Error("Error jwt token parsing : ", err.Error())
		err := errors.InternalError.Wrap(err, "Error jwt token parsing")
		responses.HandleError(c, err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		clients, err := cc.clientService.WithTrx(trx).LoginClient(claims["email"].(string))
		if err != nil {
			err := errors.Conflict.New("Email not found")
			errs := errors.SetCustomMessage(err, "Email not found")
			responses.HandleError(c, errs)
			return
		}
		if claims["email"].(string) == clients.Email {

			access_token, err, refresh_token, refresherror := utils.GenerateJWT(claims["email"].(string))

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

}
