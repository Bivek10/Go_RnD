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

type HistoryController struct {
	logger              infrastructure.Logger
	quizHistoryServices services.QuizHistoryServices
	env                 infrastructure.Env
	firbaseServices     services.FirebaseService
}

func NewHistoryController(logger infrastructure.Logger, quizHistoryService services.QuizHistoryServices, friebaseService services.FirebaseService, evn infrastructure.Env) HistoryController {
	return HistoryController{
		logger:              logger,
		env:                 evn,
		firbaseServices:     friebaseService,
		quizHistoryServices: quizHistoryService,
	}

}

// create ->quiz
func (qq HistoryController) CreateHistory(c *gin.Context) {
	quizhistory := models.QuizHistory{}
	trx := c.MustGet(constants.DBTransaction).(*gorm.DB)
	if err := c.ShouldBindJSON(&quizhistory); err != nil {
		qq.logger.Zap.Error("Error [create History] (should Bind Json): ", err)
		responses.HandleError(c, err)
		return
	}

	if err := qq.quizHistoryServices.WithTrx(trx).CreateHistory(quizhistory); err != nil {
		qq.logger.Zap.Error("Error [create history] [db create history]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to create history")

		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "History data created succesfully")
}

// get allquiz  data
func (qq HistoryController) GetAllHistory(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	choices, count, err := qq.quizHistoryServices.GetAllHistory(pagination)
	if err != nil {
		qq.logger.Zap.Error("Error geting history records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get history data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, choices, count)
}

// get allquiz by quiz id data
func (qq HistoryController) GetHistoryByUserID(c *gin.Context) {
	id := c.Param("user_id")
	pagination := utils.BuildPagination(c)
	user_id, errs := strconv.Atoi(id)
	if errs != nil {
		qq.logger.Zap.Error("Error converting the string into int", errs.Error())
		err := errors.InternalError.Wrap(errs, "Failed failed to convert error to int")
		responses.HandleError(c, err)
		return
	}
	quizhistory, count, err := qq.quizHistoryServices.GetHistoryByUserID(pagination, user_id)
	if err != nil {
		qq.logger.Zap.Error("Error geting quiz history records", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to get quiz history data")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, quizhistory, count)
}
