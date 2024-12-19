package handlers

import (
	authService "notify-backend/api/auth/services"
	apiResponse "notify-backend/api/utils/response"

	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	authService.LoginService()
	apiResponse.SuccessResponse("Login handler", nil)

	c.JSON(http.StatusOK, apiResponse.SuccessResponse("Login handler", nil))
}
