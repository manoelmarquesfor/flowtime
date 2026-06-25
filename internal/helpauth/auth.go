package helpauth

import (
	"context"
	"net/http"

	"github/manoelmarquesfor/flowtime/internal/errs"
)

type contextKey string

const (
	userContextKey contextKey = "usuario"
)

func GetUserRequisicao(r *http.Request) (UsuarioAutenticado, error) {
	userContext := r.Context().Value(userContextKey)
	if userContext == nil {
		return UsuarioAutenticado{}, errs.NewInternalError("Usuário não encontrado no contexto da requisição")
	}

	usuario, ok := userContext.(*UsuarioAutenticado)
	if !ok {
		return UsuarioAutenticado{}, errs.NewInternalError("Usuário no contexto da requisição tem tipo incorreto")
	}

	return *usuario, nil
}

func SetUserRequisicao(r *http.Request, user *UsuarioAutenticado) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), userContextKey, user))
}
