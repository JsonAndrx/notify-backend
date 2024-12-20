package routes

import (
	authHandlers "notify-backend/api/auth/handlers"

	"github.com/gin-gonic/gin"
)

func SeptupRoutes(router *gin.Engine) {
	routePrefix := router.Group("/api/v1/")

	authGroup := routePrefix.Group("/auth")
	{
		authGroup.POST("/login", authHandlers.LoginHandler)
		authGroup.POST("/register", authHandlers.RegisterHandler)

		contentGroup := authGroup.Group("/content")
		{
			contentGroup.GET("/list_countries", authHandlers.ListCountriesHandler)
			contentGroup.POST("/get_timezone", authHandlers.GetTimezoneByCountryId)
		}
	}

}
