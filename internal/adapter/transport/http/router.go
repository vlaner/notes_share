package http_transport

import (
	"log"
	"net/http"
	"strings"

	"github.com/vlaner/notes_share/pkg/html"
)

type Router struct {
	htmlRenderer *html.Renderer
	noteHandler  *NoteHandler
}

func NewHandler(htmlRenderer *html.Renderer, noteHandler *NoteHandler) *Router {
	return &Router{
		htmlRenderer: htmlRenderer,
		noteHandler:  noteHandler,
	}
}

func stopDirectoryListing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (r *Router) RegisterRoutes(staticDir string, mux *http.ServeMux) {
	fs := http.FileServer(http.Dir(staticDir))

	mux.Handle("GET /static/", http.StripPrefix("/static", stopDirectoryListing(fs)))

	mux.HandleFunc("GET /{$}", r.handleIndex)

	mux.HandleFunc("POST /api/notes", r.noteHandler.CreateNote)
	mux.HandleFunc("GET /notes/{id}", r.noteHandler.GetNote)

}

func (h *Router) handleIndex(w http.ResponseWriter, r *http.Request) {
	err := h.htmlRenderer.Render(w, "index", nil)
	if err != nil {
		log.Println("error rendering", err)
	}
}
