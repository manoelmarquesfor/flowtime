package usuario

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
	router.Route("/usuario", func(r chi.Router) {
		r.Use(validateSessaoMiddleware.ValidateCookie)
		r.Get("/", h.getAll)
		r.Get("/{id}", h.get)
		r.Post("/", h.create)
		r.Delete("/{id}", h.delete)
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

	users, err := h.service.getAll(r.Context(), usuarioAutenticado)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	webutil.Response(w, http.StatusOK, users)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	usuarioAutenticado, err := helpauth.GetUserRequisicao(r)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	var usuario UsuarioCreate

	err = webutil.DecodeJSON(r.Body, &usuario)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	userCreated, err := h.service.create(r.Context(), usuario, usuarioAutenticado)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	webutil.Response(w, http.StatusCreated, userCreated)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	usuarioAutenticado, err := helpauth.GetUserRequisicao(r)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	id := chi.URLParam(r, "id")

	err = h.service.delete(r.Context(), id, usuarioAutenticado)
	if err != nil {
		webutil.ResponseError(w, err)

		return
	}

	response := DeleteUsuarioResponse{
		ID:     id,
		Detail: "Usuário deletado com sucesso",
	}

	webutil.Response(w, http.StatusOK, response)
}
