package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const (
	MAXAGE int = 300
)

func Setup(router *chi.Mux) {
	// Configuração do middleware de CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // domínio do site
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           MAXAGE, // tempo em segundos
	})

	router.Use(middleware.Logger)
	router.Use(corsMiddleware.Handler)
}
