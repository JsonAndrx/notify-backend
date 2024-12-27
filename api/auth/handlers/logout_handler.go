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
    http.SetCookie(c.Writer, &http.Cookie{
        Name:     "token",
        Value:    "",
        Expires:  time.Now().Add(-1 * time.Hour),
        HttpOnly: true,
        SameSite: http.SameSiteNoneMode,
        Secure:   true,
    })

    c.JSON(http.StatusOK, apiResponse.SuccessResponse("Logout successful", nil))
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, apiResponse.SuccessResponse("Test successful", nil))
}