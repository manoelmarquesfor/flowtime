package colaborador

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
	router.Route("/colaborador", func(r chi.Router) {
		r.Use(validateSessaoMiddleware.ValidateCookie)
		r.Get("/", h.getAll)
		r.Get("/{id}", h.get)
		r.Post("/", h.create)
		r.Put("/alterar-situacao/{id}", h.alterarSituacao)
	})
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	usuarioAutenticado, err := helpauth.GetUserRequisicao(r)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	id := chi.URLParam(r, "id")

	user, err := h.service.get(r.Context(), id, usuarioAutenticado)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	webutil.Response(w, http.StatusOK, user)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	usuarioAutenticado, err := helpauth.GetUserRequisicao(r)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	ativo := r.URL.Query().Get("ativo")

	colaboradores, err := h.service.getAll(r.Context(), usuarioAutenticado, ativo)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	webutil.Response(w, http.StatusOK, colaboradores)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	usuarioAutenticado, err := helpauth.GetUserRequisicao(r)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	var colaborador CreateRequest

	err = webutil.DecodeJSON(r.Body, &colaborador)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	colaboradorCreated, err := h.service.create(r.Context(), colaborador, usuarioAutenticado)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	webutil.Response(w, http.StatusCreated, colaboradorCreated)
}

func (h *Handler) alterarSituacao(w http.ResponseWriter, r *http.Request) {
	usuarioAutenticado, err := helpauth.GetUserRequisicao(r)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	id := chi.URLParam(r, "id")

	colaboradorUpdated, err := h.service.alterarSituacao(r.Context(), id, usuarioAutenticado)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	webutil.Response(w, http.StatusOK, colaboradorUpdated)
}
