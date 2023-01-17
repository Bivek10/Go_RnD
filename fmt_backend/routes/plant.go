package routes

import (
	"github.com/bivek/fmt_backend/controllers"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/middlewares"
)

//plantRoutes -> struct

type PlantRoutes struct {
	logger          infrastructure.Logger
	router          infrastructure.Router
	plantController controllers.PlantController
	middlewares     middlewares.FirebaseAuthMiddleware
	trxMiddleware   middlewares.DBTransactionMiddleware
}

// setup plant routes
func (i PlantRoutes) Setup() {

	i.logger.Zap.Info("setting up plants routes")
	plants := i.router.Gin.Group("/plants")
	{

		plants.GET("",i.middlewares.Handle() ,i.plantController.GetAllPlant)
		plants.POST("addPlant", i.trxMiddleware.DBTransactionHandle(), i.plantController.CreatePlant)
	}
}

// NewPlantRoutes -> creates new plant
func NewPlantRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	plantController controllers.PlantController,
	middlerware middlewares.FirebaseAuthMiddleware,
	trxMiddleware middlewares.DBTransactionMiddleware,
) PlantRoutes {
	return PlantRoutes{
		logger:          logger,
		router:          router,
		plantController: plantController,
		middlewares:     middlerware,
		trxMiddleware:   trxMiddleware,
	}
}
