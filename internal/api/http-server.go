package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/vandaimer/workouts/config"
)

type Server struct {
	server *http.Server
	config *config.AppConfig
}

func NewHTTPServer(config *config.AppConfig, router *gin.Engine) *Server {
	return &Server{
		server: &http.Server{Addr: ":" + config.Port, Handler: router},
		config: config,
	}
}

func (server *Server) Run() {
	log.Info().Str("port", server.config.Port).Msg("HTTP Server running.")
	if err := server.server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}
