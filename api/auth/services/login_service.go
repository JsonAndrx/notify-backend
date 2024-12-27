package services

import (
    "net/http"
    "notify-backend/api/auth/repositories"
    "notify-backend/api/auth/types"
    "notify-backend/api/utils/encrypt"
    "golang.org/x/crypto/bcrypt"
)

func LoginService(loginRequest types.LoginRequest, w http.ResponseWriter) (bool, error) {
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

    token, expirationTime, err := encrypt.GenerateToken(user.Username)
    if err != nil {
        return false, err
    }

    http.SetCookie(w, &http.Cookie{
        Name:     "token",
        Value:    token,
        Expires:  expirationTime,
        SameSite: http.SameSiteNoneMode,
        HttpOnly: true,
        Secure:   true,
    })

    return true, nil
}