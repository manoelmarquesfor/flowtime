package middleware

import (
	"errors"
	"net/http"

	"github/manoelmarquesfor/flowtime/internal/errs"
	"github/manoelmarquesfor/flowtime/internal/helpauth"

	"github/manoelmarquesfor/flowtime/internal/webutil"
)

type ValidateSessaoMiddleware struct {
	helpAuthRepository *helpauth.Repository
}

func NewValidateSessaoMiddleware(helpAuthRepository *helpauth.Repository) *ValidateSessaoMiddleware {
	return &ValidateSessaoMiddleware{
		helpAuthRepository: helpAuthRepository,
	}
}

func (m *ValidateSessaoMiddleware) ValidateCookie(h http.Handler) http.Handler {
	fn := func(writer http.ResponseWriter, request *http.Request) {
		sessao := webutil.GetSessionCookie(request)
		if sessao == nil {
			webutil.ResponseError(writer, errs.NewUnauthorizedError())

			return
		}

		usuarioRepository, err := m.helpAuthRepository.UsuarioBySessao(request.Context(), *sessao)
		if err != nil {
			var notFoundErr *errs.NotFoundError
			if errors.As(err, &notFoundErr) {
				webutil.ResponseError(writer, errs.NewUnauthorizedError())

				return
			}

			webutil.ResponseError(writer, err)

			return
		}

		h.ServeHTTP(writer, helpauth.SetUserRequisicao(request, usuarioRepository.ToUsuarioAutenticado()))
	}

	return http.HandlerFunc(fn)
}
