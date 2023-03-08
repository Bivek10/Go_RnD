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
	jwtMiddleware     middlewares.JWTAuthMiddleWare
	trxMiddleware     middlewares.DBTransactionMiddleware
}

// Setup user routes
func (i ClientRoutes) Setup() {
	i.logger.Zap.Info(" Setting up client routes")
	newusers := i.router.Gin.Group("/")
	{
		//users.GET("", i.ClientsController.GetAllUsers)
		newusers.POST("login", i.trxMiddleware.DBTransactionHandle(), i.clientsController.LoginClient)
		newusers.POST("register", i.trxMiddleware.DBTransactionHandle(), i.clientsController.CreateClients)
		newusers.POST("refreshToken", i.clientsController.ReGenerateClientToken)
		newusers.Static("files", "./clients-image")
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
	jwtMiddlewware middlewares.JWTAuthMiddleWare,
) ClientRoutes {
	return ClientRoutes{
		router:            router,
		logger:            logger,
		clientsController: clientsController,
		middleware:        middleware,
		trxMiddleware:     trxMiddleware,
		jwtMiddleware:     jwtMiddlewware,
	}
}
