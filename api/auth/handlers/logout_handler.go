package handlers

import (
	"net/http"
	apiResponse "notify-backend/api/utils/response"
	"time"
    "fmt"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
    cookie, err := c.Request.Cookie("token")
    if err != nil {
        if err == http.ErrNoCookie {
            c.JSON(http.StatusUnauthorized, apiResponse.ErrorResponse("No active session", nil))
            return
        }
        c.JSON(http.StatusInternalServerError, apiResponse.ErrorResponse("Error retrieving session", err))
        return
    }

    fmt.Println(cookie)
    // Si la cookie existe, eliminarla estableciendo una fecha de expiración en el pasado
    c.SetSameSite(http.SameSiteNoneMode)
    c.SetCookie("token", "", -1, "/", c.Request.Host, true, true)

    c.JSON(http.StatusOK, apiResponse.SuccessResponse("Logout successful", nil))
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, apiResponse.SuccessResponse("Test successful", nil))
}