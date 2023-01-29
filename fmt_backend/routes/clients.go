package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/middlewares"
)

// ClientRoutes -> struct
type ClientRoutes struct {
	logger            infrastructure.Logger
	router            infrastructure.Router
	clientsController controllers.ClientController
	middleware        middlewares.FirebaseAuthMiddleware
	trxMiddleware     middlewares.DBTransactionMiddleware
}

// Setup user routes
func (i ClientRoutes) Setup() {
	i.logger.Zap.Info(" Setting up client routes")
	newusers := i.router.Gin.Group("/newuser")
	{
		//users.GET("", i.ClientsController.GetAllUsers)
		newusers.POST("create", i.trxMiddleware.DBTransactionHandle(), i.clientsController.CreateClients)
		//users.POST("login", i.ClientsController.UserLogin)
	}
}

// NewClientRoutes -> creates new user controller
func NewClientRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	clientsController controllers.ClientController,
	middleware middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
) ClientRoutes {
	return ClientRoutes{
		router:            router,
		logger:            logger,
		clientsController: clientsController,
		middleware:        middleware,
		trxMiddleware:     trxMiddleware,
	}
}
