package handlers

import (
	"net/http"
	authService "notify-backend/api/auth/services"
	apiResponse "notify-backend/api/utils/response"

	"github.com/gin-gonic/gin"
)

func ListCountriesHandler(c *gin.Context) {
	listCountries, err := authService.ListCountriesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiResponse.ErrorResponse("Error while fetching countries", err))
		return
	}

	c.JSON(http.StatusOK, apiResponse.SuccessResponse("List countries fetching", listCountries))
}
