package controllers

import (
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/services"
)

type QuestionsController struct {
	logger          infrastructure.Logger
	questionService services.QuestionServices
	env             infrastructure.Env
	firbaseServices services.FirebaseService
}

func NewQuestionController(logger infrastructure.Logger, questionService services.QuestionServices, friebaseService services.FirebaseService, evn infrastructure.Env) QuestionsController {
	return QuestionsController{
		logger:          logger,
		env:             evn,
		firbaseServices: friebaseService,
		questionService: questionService,
	}

}
