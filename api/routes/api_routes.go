package routes

import (
	"github.com/gin-gonic/gin"
	authHandlers "notify-backend/api/auth/handlers"
)

func SeptupRoutes(router *gin.Engine) {
	routePrefix := router.Group("/api/v1/")

	authGroup := routePrefix.Group("/auth")
	{
		authGroup.POST("/login", authHandlers.LoginHandler)
		
		contentGroup := authGroup.Group("/content")
		{
			contentGroup.GET("/list_countries", authHandlers.ListCountriesHandler)
			contentGroup.POST("/get_timezone", authHandlers.GetTimezoneByCountryId)
		}
	}


}