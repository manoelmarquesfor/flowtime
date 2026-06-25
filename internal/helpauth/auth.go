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

func GetUserRequisicao(r *http.Request) (UsuarioAutencicado, error) {
	userContext := r.Context().Value(userContextKey)
	if userContext == nil {
		return UsuarioAutencicado{}, errs.NewInternalError("Usuário não encontrado no contexto da requisição")
	}

	usuario, ok := userContext.(*UsuarioAutencicado)
	if !ok {
		return UsuarioAutencicado{}, errs.NewInternalError("Usuário no contexto da requisição tem tipo incorreto")
	}

	return *usuario, nil
}

func SetUserRequisicao(r *http.Request, user *UsuarioAutencicado) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), userContextKey, user))
}
