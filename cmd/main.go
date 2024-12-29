package main

import (
    "fmt"
    "net/http"
    "notify-backend/api/routes"
    "notify-backend/api/utils/debug"
    "notify-backend/api/utils/response"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/rs/zerolog/log"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        debug.LogError(err)
    }

    log.Info().Msg("Starting server")
    r := gin.Default()

    // Configurar CORS para permitir solicitudes desde dominios específicos
    r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
            return true
        },
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    fmt.Println("paso los cors")
    routes.SeptupRoutes(r)

    server := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }

    servErr := server.ListenAndServe()
    response.ErrorResponse("Error starting server", servErr)
}