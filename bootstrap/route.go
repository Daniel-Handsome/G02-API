package bootstrap

import (
	middlewares "G02-API/app/http/middwares"
	"G02-API/routes"
	"github.com/gin-gonic/gin"
)

func SetUpRoute(route *gin.Engine) {
	// globaMiddware
	registerGlobalMiddwate(route)

	// register api route
	routes.RegisterApiRoutes(route)

	// register 404 page
	routes.Register404Handler(route)
}

func registerGlobalMiddwate(r *gin.Engine) {
	r.Use(middlewares.Logger(),
		middlewares.Recovery(),
	)
}
