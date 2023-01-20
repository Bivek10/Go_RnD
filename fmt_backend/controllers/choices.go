package controllers

import (
	"net/http"
	"strconv"

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

type ChoiceController struct {
	logger          infrastructure.Logger
	choiceService   services.ChoiceServices
	env             infrastructure.Env
	firbaseServices services.FirebaseService
}

func NewChoiceController(logger infrastructure.Logger, choiceService services.ChoiceServices, friebaseService services.FirebaseService, evn infrastructure.Env) ChoiceController {
	return ChoiceController{
		logger:          logger,
		env:             evn,
		firbaseServices: friebaseService,
		choiceService:   choiceService,
	}

}

// create ->quiz
func (qq ChoiceController) CreateChoices(c *gin.Context) {
	choices := models.Choices{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)
	if err := c.ShouldBindJSON(&choices); err != nil {
		qq.logger.Zap.Error("Error [create choice] (should Bind Json): ", err)
		responses.HandleError(c, err)
		return
	}

	if err := qq.choiceService.WithTrx(trx).CreateChoices(choices); err != nil {
		qq.logger.Zap.Error("Error [create user] [db create choices]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create choices")

		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Choices data created succesfully")
}

// get allquiz  data
func (qq ChoiceController) GetAllChoices(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	choices, count, err := qq.choiceService.GetAllChoices(pagination)
	if err != nil {
		qq.logger.Zap.Error("Error geting choices records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get choices data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, choices, count)
}

// get allquiz by quiz id data
func (qq ChoiceController) GetChoicesByQuestionID(c *gin.Context) {
	id := c.Param("question_id")
	pagination := utils.BuildPagination(c)
	quiz_id, errs := strconv.Atoi(id)
	if errs != nil {
		qq.logger.Zap.Error("Error converting the string into int", errs.Error())
		err := errors.InternalError.Wrap(errs, "Failed failed to convert error to int")
		responses.HandleError(c, err)
		return
	}
	choice, count, err := qq.choiceService.GetChoicesByQuestionID(pagination, int64(quiz_id))
	if err != nil {
		qq.logger.Zap.Error("Error geting choices records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get choices data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, choice, count)
}
