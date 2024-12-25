package routes

import (
	authHandlers "notify-backend/api/auth/handlers"
	"notify-backend/api/middlewares"

	"github.com/gin-gonic/gin"
)

func SeptupRoutes(router *gin.Engine) {
	routePrefix := router.Group("/api/v1/")

	authGroup := routePrefix.Group("/auth")
	{
		authGroup.POST("/login", authHandlers.LoginHandler)
		authGroup.POST("/register", authHandlers.RegisterHandler)
		authGroup.POST("/logout", authHandlers.LogoutHandler)
		authGroup.GET("/verify-session", authHandlers.VerifySessionHandler)

		contentGroup := authGroup.Group("/content")
		{
			contentGroup.GET("/list_countries", authHandlers.ListCountriesHandler)
			contentGroup.POST("/get_timezone", authHandlers.GetTimezoneByCountryId)
		}

		dashboardGroup := authGroup.Group("/dashboard")
		dashboardGroup.Use(middlewares.AuthMiddleware())
		{
			dashboardGroup.GET("/test", authHandlers.Test)
		}

	}

}
