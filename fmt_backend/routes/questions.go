package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/middlewares"
)

type QuestionsRoutes struct {
	logger             infrastructure.Logger
	router             infrastructure.Router
	questionController controllers.QuestionsController
	middlewares        middlewares.FirebaseAuthMiddleware
	trxMiddleware      middlewares.DBTransactionMiddleware
	jwtMiddleware      middlewares.JWTAuthMiddleWare
}

//setup quiz routes

func (i QuestionsRoutes) Setup() {
	i.logger.Zap.Info("setting up questions routes")
	quizs := i.router.Gin.Group("/questions")
	{
		quizs.GET("", i.jwtMiddleware.Handle(), i.questionController.GetAllQuestion)
		quizs.GET(":quiz_id", i.jwtMiddleware.Handle(), i.questionController.GetQuestionByID)

		quizs.POST("create", i.trxMiddleware.DBTransactionHandle(), i.questionController.CreateQuestion)
	}
}

//NewQuizRoutes -> creates new quiz

func NewQuestionRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	question controllers.QuestionsController,
	middlewares middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
	jwtMiddlware middlewares.JWTAuthMiddleWare,
) QuestionsRoutes {
	return QuestionsRoutes{
		logger:             logger,
		router:             router,
		questionController: question,
		middlewares:        middlewares,
		trxMiddleware:      trxMiddleware,
		jwtMiddleware:      jwtMiddlware,
	}
}
