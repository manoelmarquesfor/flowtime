package feriado

import (
	"net/http"

	"github/manoelmarquesfor/flowtime/internal/helpauth"
	"github/manoelmarquesfor/flowtime/internal/middleware"
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

func (h *Handler) RegisterRoutes(router chi.Router, validateSessaoMiddleware *middleware.ValidateSessaoMiddleware) {
	router.Route("/feriado", func(r chi.Router) {
		r.Use(validateSessaoMiddleware.ValidateCookie)
		r.Post("/", h.create)
		r.Get("/", h.all)
		r.Delete("/{data}", h.delete)
	})
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	usuarioAutenticado, err := helpauth.GetUserRequisicao(r)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	var feriadoRequest CreateRequest

	err = webutil.DecodeJSON(r.Body, &feriadoRequest)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	err = h.service.CreateFeriado(r.Context(), feriadoRequest, usuarioAutenticado)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	webutil.Response(w, http.StatusOK, CreateResponse{Detail: "Feriado criado com sucesso"})
}

func (h *Handler) all(w http.ResponseWriter, r *http.Request) {
	ano := r.URL.Query().Get("ano")

	feriados, err := h.service.GetFeriados(r.Context(), ano)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	webutil.Response(w, http.StatusOK, feriados)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "data")

	usuarioAutenticado, err := helpauth.GetUserRequisicao(r)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	err = h.service.DeleteFeriado(r.Context(), data, usuarioAutenticado)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	webutil.Response(w, http.StatusOK, DeleteResponse{Detail: "Feriado deletado com sucesso"})
}
