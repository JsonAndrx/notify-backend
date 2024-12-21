package middlewares

import (
	"fmt"
	"net/http"
	"notify-backend/api/utils/encrypt"
    apiResponse "notify-backend/api/utils/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("Auth middleware")
        cookie, err := c.Request.Cookie("token")
        if err != nil {
            if err == http.ErrNoCookie {
                c.JSON(http.StatusUnauthorized, apiResponse.ErrorResponse("Unauthorized", err))
                c.Abort()
                return
            }
            c.JSON(http.StatusUnauthorized, apiResponse.ErrorResponse("Unauthorized", err))
            c.Abort()
            return
        }

        tokenString := cookie.Value
        claims, valid := encrypt.ValidateToken(tokenString)
        if !valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("user", claims.Username)
        c.Next()
    }
}