package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed pages/*.html assets/*
var embeddedFiles embed.FS

var (
	pagesFS, _  = fs.Sub(embeddedFiles, "pages")
	assetsFS, _ = fs.Sub(embeddedFiles, "assets")
)

func Setup(router *chi.Mux) {
	router.Get("/", servePage("index.html"))
	router.Get("/login", servePage("login.html"))
	router.Get("/colaboradores", servePage("colaboradores.html"))
	router.Get("/feriado", servePage("feriado.html"))
	router.Get("/relatorio", servePage("relatorio.html"))
	router.Get("/ponto-matricula", servePage("ponto-matricula.html"))
	router.Get("/ponto-tag", servePage("ponto-tag.html"))
	router.Get("/usuario", servePage("usuario.html"))

	router.Handle("/assets/*",
		http.StripPrefix("/assets/",
			http.FileServer(http.FS(assetsFS)),
		),
	)
}

func servePage(file string) http.HandlerFunc {
	return func(writer http.ResponseWriter, response *http.Request) {
		data, err := fs.ReadFile(pagesFS, file)
		if err != nil {
			http.NotFound(writer, response)

			return
		}

		writer.Header().Set("Content-Type", "text/html")

		_, err = writer.Write(data)
		if err != nil {
			http.Error(writer, "Failed to write response", http.StatusInternalServerError)
		}
	}
}
