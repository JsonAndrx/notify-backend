package main

import (
	"fmt"
	"net/http"
	"notify-backend/api/routes"
	response "notify-backend/api/utils/response"

	"notify-backend/api/utils/debug"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		debug.LogError(err)
	}

	log.Info().Msg("Starting server")
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:    []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "access-control-allow-origin", "access-control-allow-headers"},
		ExposeHeaders:    []string{"Set-Cookie"}, // Exponer encabezados específicos si es necesario
		AllowCredentials: true, // Permite cookies y credenciales
		MaxAge:       12 * time.Hour,
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
