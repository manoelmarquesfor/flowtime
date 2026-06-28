package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github/manoelmarquesfor/flowtime/internal/auth"
	"github/manoelmarquesfor/flowtime/internal/colaborador"
	"github/manoelmarquesfor/flowtime/internal/config"
	"github/manoelmarquesfor/flowtime/internal/database"
	"github/manoelmarquesfor/flowtime/internal/feriado"
	"github/manoelmarquesfor/flowtime/internal/helpauth"
	"github/manoelmarquesfor/flowtime/internal/middleware"
	"github/manoelmarquesfor/flowtime/internal/usuario"
	"github/manoelmarquesfor/flowtime/web"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

const (
	TEMP string = "./temp"

	HEADERTIMEOUT time.Duration = 10 * time.Second
	PERMISSAO     os.FileMode   = 0o755
)

func carregarInit() {
	err := os.MkdirAll(TEMP, PERMISSAO)
	if errors.Is(err, os.ErrExist) {
		log.Println("Diretório temporário já existe.")
	}
}

func main() {
	log.Println("Iniciando o servidor...")
	carregarInit()

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	srv := setupServer(config)

	go func() {
		log.Printf("Server is running on port %d \n", config.Server.Port)

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Server.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Print(err.Error())

		return
	}

	<-ctx.Done()
	log.Println("Server shutdown")

	log.Println("server exiting")
}

func setupServer(config *config.Config) *http.Server {
	database, err := database.Setup(config)
	if err != nil {
		log.Fatal(err)
	}

	helpauthRepository := helpauth.NewRepository(database)
	validateSessaoMiddleware := middleware.NewValidateSessaoMiddleware(helpauthRepository)

	router := chi.NewRouter()
	middleware.Setup(router)
	web.Setup(router)

	router.Route("/api", func(api chi.Router) {
		setupRouterAuth(api, database)
		setupRouterUsuario(api, database, validateSessaoMiddleware)
		setupRouterFeriado(api, database, validateSessaoMiddleware)
		setupRouterColaborador(api, database, validateSessaoMiddleware)
	})
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.Server.Port),
		Handler:           router,
		ReadHeaderTimeout: HEADERTIMEOUT,
		WriteTimeout:      HEADERTIMEOUT,
	}

	return srv
}

func setupRouterAuth(router chi.Router, database *sqlx.DB) {
	authRepository := auth.NewRepository(database)
	authService := auth.NewService(authRepository)
	authHandler := auth.NewHandler(authService)

	authHandler.RegisterRoutes(router)
}

func setupRouterUsuario(
	router chi.Router,
	database *sqlx.DB,
	validateSessaoMiddleware *middleware.ValidateSessaoMiddleware,
) {
	usuarioRepository := usuario.NewRepository(database)
	usuarioService := usuario.NewService(usuarioRepository)
	usuarioHandler := usuario.NewHandler(usuarioService)

	usuarioHandler.RegisterRoutes(router, validateSessaoMiddleware)
}

func setupRouterFeriado(
	router chi.Router,
	database *sqlx.DB,
	validateSessaoMiddleware *middleware.ValidateSessaoMiddleware,
) {
	feriadoRepository := feriado.NewRepository(database)
	feriadoService := feriado.NewService(feriadoRepository)
	feriadoHandler := feriado.NewHandler(feriadoService)

	feriadoHandler.RegisterRoutes(router, validateSessaoMiddleware)
}

func setupRouterColaborador(
	router chi.Router,
	database *sqlx.DB,
	validateSessaoMiddleware *middleware.ValidateSessaoMiddleware,
) {
	colaboradorRepository := colaborador.NewRepository(database)
	colaboradorService := colaborador.NewService(colaboradorRepository)
	colaboradorHandler := colaborador.NewHandler(colaboradorService)

	colaboradorHandler.RegisterRoutes(router, validateSessaoMiddleware)
}
