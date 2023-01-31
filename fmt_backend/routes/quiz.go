package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/middlewares"
)

type QuizRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	quizController controllers.QuizController
	middlewares    middlewares.FirebaseAuthMiddleware
	trxMiddleware  middlewares.DBTransactionMiddleware
	jwtMiddleware  middlewares.JWTAuthMiddleWare
}

//setup quiz routes

func (i QuizRoutes) Setup() {
	i.logger.Zap.Info("setting up quiz routes")
	quizs := i.router.Gin.Group("/quizs")
	{
		quizs.GET("", i.jwtMiddleware.Handle(), i.quizController.GetAllQuiz)
		quizs.POST("create", i.trxMiddleware.DBTransactionHandle(), i.quizController.CreateQuiz)
	}
}

//NewQuizRoutes -> creates new quiz

func NewQuizRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	quizController controllers.QuizController,
	middlewares middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
	jwtMiddleware middlewares.JWTAuthMiddleWare,
) QuizRoutes {
	return QuizRoutes{
		logger:         logger,
		router:         router,
		quizController: quizController,
		middlewares:    middlewares,
		trxMiddleware:  trxMiddleware,
		jwtMiddleware:  jwtMiddleware,
	}
}
