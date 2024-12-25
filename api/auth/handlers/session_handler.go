package handlers

import (
    "net/http"
    apiResponse "notify-backend/api/utils/response"
    "notify-backend/api/utils/encrypt"
    "github.com/gin-gonic/gin"
)

func VerifySessionHandler(c *gin.Context) {
    cookie, err := c.Request.Cookie("token")
    if err != nil {
        if err == http.ErrNoCookie {
            c.JSON(http.StatusUnauthorized, apiResponse.ErrorResponse("No active session", nil))
            return
        }
        c.JSON(http.StatusInternalServerError, apiResponse.ErrorResponse("Error retrieving session", err))
        return
    }

    tokenString := cookie.Value
    claims, valid := encrypt.ValidateToken(tokenString)
    if !valid {
        c.JSON(http.StatusUnauthorized, apiResponse.ErrorResponse("Invalid token", nil))
        return
    }

    c.JSON(http.StatusOK, apiResponse.SuccessResponse("Session is valid", gin.H{
        "user": claims.Username,
    }))
}