package services

import (
    "net/http"
    "notify-backend/api/auth/repositories"
    "notify-backend/api/auth/types"
    "notify-backend/api/utils/encrypt"
    "golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
)

func LoginService(loginRequest types.LoginRequest, c *gin.Context) (bool, error) {
    user, err := repositories.GetUserByUsernameOrEmail(loginRequest.Username)
    if err != nil {
        return false, err
    }

    if user == (types.User{}) {
        return false, nil
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
    if err != nil {
        return false, err
    }

    token, _, err := encrypt.GenerateToken(user.Username)
    if err != nil {
        return false, err
    }

    c.SetSameSite(http.SameSiteNoneMode)
    c.SetCookie("token", token, 3600*24*30, "", "", false, true)

    return true, nil
}