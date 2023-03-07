package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/middlewares"
)

type HistoryRoutes struct {
	logger            infrastructure.Logger
	router            infrastructure.Router
	historyController controllers.HistoryController
	middlewares       middlewares.FirebaseAuthMiddleware
	trxMiddleware     middlewares.DBTransactionMiddleware
	jwtMiddleware     middlewares.JWTAuthMiddleWare
}

//setup quiz routes

func (i HistoryRoutes) Setup() {
	i.logger.Zap.Info("setting up Histroy routes")
	quizs := i.router.Gin.Group("/history")
	{
		quizs.GET("", i.historyController.GetAllHistory)
		quizs.GET(":user_id", i.jwtMiddleware.Handle(), i.historyController.GetHistoryByUserID)

		quizs.POST("", i.jwtMiddleware.Handle(), i.trxMiddleware.DBTransactionHandle(), i.historyController.CreateHistory)
	}
}

//NewQuizRoutes -> creates new quiz

func NewHistoryRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	historyController controllers.HistoryController,
	middlewares middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
	jwtMiddlerware middlewares.JWTAuthMiddleWare,
) HistoryRoutes {
	return HistoryRoutes{
		logger:            logger,
		router:            router,
		historyController: historyController,
		middlewares:       middlewares,
		trxMiddleware:     trxMiddleware,
		jwtMiddleware:     jwtMiddlerware,
	}
}
