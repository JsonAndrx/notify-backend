package handlers

import (
	"net/http"
	"notify-backend/api/auth/types"
	"notify-backend/api/utils/debug"
	"notify-backend/api/auth/services"
	apiResponse "notify-backend/api/utils/response"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var registerRequest types.RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		debug.LogError(err)
		c.JSON(http.StatusBadRequest, apiResponse.ErrorResponse(err.Error(), nil))
		return
	}

	validRegister, message, err  := services.RegisterService(registerRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiResponse.ErrorResponse("", nil))
		return
	}

	if !validRegister {
		c.JSON(http.StatusBadRequest, apiResponse.ErrorResponse(message, nil))
		return
	}

	c.JSON(http.StatusOK, apiResponse.SuccessResponse(message, nil))

}
