package main

import (
	"net/http"
	response "notify-backend/api/utils/response"
	"notify-backend/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting server")
	r := gin.Default()
	routes.SeptupRoutes(r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	servErr := server.ListenAndServe()
	response.ErrorResponse("Error starting server", servErr)
}
