package handlers

import (
	authService "notify-backend/api/auth/services"
	apiResponse "notify-backend/api/utils/response"
	"notify-backend/api/auth/types"

	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {
    var loginRequest types.LoginRequest
    if err := c.ShouldBindJSON(&loginRequest); err != nil {
        c.JSON(http.StatusBadRequest, apiResponse.ErrorResponse("Invalid request", err))
        return
    }

    validLogin, err := authService.LoginService(loginRequest, c.Writer)
    if err != nil {
        c.JSON(http.StatusUnauthorized, apiResponse.ErrorResponse("Invalid username or password", err))
        return
    }

	if !validLogin {
		c.JSON(http.StatusUnauthorized, apiResponse.ErrorResponse("Invalid username or password", nil))
		return
	}

    c.JSON(http.StatusOK, apiResponse.SuccessResponse("Login successful", nil))
}