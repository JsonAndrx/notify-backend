package handlers

import (
	"net/http"
	authService "notify-backend/api/auth/services"
	apiResponse "notify-backend/api/utils/response"
	authTypes "notify-backend/api/auth/types"

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

func GetTimezoneByCountryId(c *gin.Context) {
	var countryId authTypes.GetTimezoneByCountryIdRequest
	if err := c.ShouldBindJSON(&countryId); err != nil {
		c.JSON(http.StatusBadRequest, apiResponse.ErrorResponse("", err))
		return
	}


	timeZones, err := authService.GetTimeZoneService(countryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiResponse.ErrorResponse("Error while fetching timezones", err))
		return
	}


	c.JSON(http.StatusOK, apiResponse.SuccessResponse("List timezones fetching", timeZones))
} 
