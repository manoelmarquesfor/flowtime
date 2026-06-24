package auth

import (
	"net/http"

	"github/manoelmarquesfor/flowtime/internal/errs"
	"github/manoelmarquesfor/flowtime/internal/webutil"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(router chi.Router) {
	router.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Delete("/logout", h.Logout)
		r.Get("/me", h.Me)
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest

	err := webutil.DecodeJSON(r.Body, &loginRequest)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	sessionID, err := h.service.Login(r.Context(), loginRequest.Email, loginRequest.Password)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	webutil.SetSessionCookie(w, sessionID)

	loginResponse := LoginResponse{
		Detail: "Login efetuado com sucesso",
	}

	webutil.Response(w, http.StatusOK, loginResponse)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	sessionID := webutil.GetSessionCookie(r)
	if sessionID == nil {
		webutil.ResponseError(w, errs.NewUnauthorizedError())

		return
	}

	err := h.service.Logout(r.Context(), *sessionID)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	webutil.ClearSessionCookie(w)
	webutil.Response(w, http.StatusOK, LogoutResponse{Detail: "Logout efetuado com sucesso"})
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	sessionID := webutil.GetSessionCookie(r)
	if sessionID == nil {
		webutil.ResponseError(w, errs.NewUnauthorizedError())

		return
	}

	webutil.Response(w, http.StatusOK, map[string]string{"Detail": "Usuário autenticado"})
}
