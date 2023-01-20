package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/middlewares"
)

type ChoiceRoutes struct {
	logger           infrastructure.Logger
	router           infrastructure.Router
	choiceController controllers.ChoiceController
	middlewares      middlewares.FirebaseAuthMiddleware
	trxMiddleware    middlewares.DBTransactionMiddleware
}

//setup quiz routes

func (i ChoiceRoutes) Setup() {
	i.logger.Zap.Info("setting up choices routes")
	quizs := i.router.Gin.Group("/choice")
	{
		quizs.GET("", i.choiceController.GetAllChoices)
		quizs.GET(":question_id", i.choiceController.GetChoicesByQuestionID)

		quizs.POST("", i.trxMiddleware.DBTransactionHandle(), i.choiceController.CreateChoices)
	}
}

//NewQuizRoutes -> creates new quiz

func NewChoicesRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	choice controllers.ChoiceController,
	middlewares middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
) ChoiceRoutes {
	return ChoiceRoutes{
		logger:           logger,
		router:           router,
		choiceController: choice,
		middlewares:      middlewares,
		trxMiddleware:    trxMiddleware,
	}
}
