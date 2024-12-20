package types

type RegisterRequest struct {
    Company Company `json:"company" binding:"required"`
    User    User    `json:"user" binding:"required"`
}

type Company struct {
    Name       string `json:"name" binding:"required"`
    Phone      string `json:"phone" binding:"required"`
    Address    string `json:"address" binding:"required"`
    TimezoneId int    `json:"timezone_id" binding:"required"`
    CountryId  int    `json:"country_id" binding:"required"`
}

type User struct {
    Username string `json:"username" binding:"required"`
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}