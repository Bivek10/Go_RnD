package controllers

import (
	"fmt"
	"net/http"

	"github.com/bivek/fmt_backend/constants"
	"github.com/bivek/fmt_backend/errors"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/bivek/fmt_backend/responses"
	"github.com/bivek/fmt_backend/services"
	"github.com/bivek/fmt_backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController -> struct
type UserController struct {
	logger          infrastructure.Logger
	userService     services.UserService
	firebaseService services.FirebaseService
	env             infrastructure.Env
}

// NewUserController -> constructor
func NewUserController(
	logger infrastructure.Logger,
	userService services.UserService,
	firebaseSerivce services.FirebaseService,
	env infrastructure.Env,
) UserController {
	return UserController{
		logger:          logger,
		userService:     userService,
		firebaseService: firebaseSerivce,
		env:             env,
	}
}

// CreateUser -> Create User
func (cc UserController) CreateUser(c *gin.Context) {
	user := models.User{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&user); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind user data")
		responses.HandleError(c, err)
		return
	}

	userid, err := cc.firebaseService.CreateUser(user.Email, user.Password)
	fmt.Printf("")
	if err != nil {
		cc.logger.Zap.Error("Error [FirebaseUser] failed: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create firebase user")
		responses.HandleError(c, err)
		return
	}

	user.UserId = userid
	//encrypt user password.
	user.Password = utils.EncryptPassword([]byte(user.Password))

	if err := cc.userService.WithTrx(trx).CreateUser(user); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] [db CreateUser]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create user")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "User Created Sucessfully")

}

// GetAllUser -> Get All User
func (cc UserController) GetAllUsers(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	users, count, err := cc.userService.GetAllUsers(pagination)

	if err != nil {
		cc.logger.Zap.Error("Error finding user records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get users data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, users, count)
}

// userlogin
func (cc UserController) UserLogin(c *gin.Context) {
	user := models.UserLoginModel{}
	if err := c.ShouldBindJSON(&user); err != nil {
		cc.logger.Zap.Error("Error [CreateUser] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to bind user data")
		responses.HandleError(c, err)
		return
	}

	users, err := cc.firebaseService.GetUserByEmail(user.Email)

	if err != nil {
		cc.logger.Zap.Error("Email not found", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get users by email")
		responses.HandleError(c, err)
		return
	}

	fmt.Println(users.UID)
	//check from local database

	//encrypt pass
	//password := utils.EncryptPassword([]byte(user.Password))
	//fmt.Println(password)

	localusers, err := cc.userService.UserLogin(user.Email, user.Password)
	//fmt.Println(password)
	//fmt.Println(err)

	if err != nil {
		cc.logger.Zap.Error("Error: [User not found in DB]", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get user from DB")
		responses.HandleError(c, err)
	}

	responses.JSON(c, 200, localusers)

	//user, err := cc.firebaseService.GetUserByEmail(user.Email);

}
