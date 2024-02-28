package html

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Renderer struct {
	templates *template.Template
}

func NewHtmlRenderer(rootPath string) (*Renderer, error) {
	tmpl := template.Template{}

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = tmpl.ParseFiles(path)
			return err
		}

		return err
	})

	return &Renderer{&tmpl}, err
}

func (h *Renderer) Render(w io.Writer, templateName string, data any) error {
	return h.templates.ExecuteTemplate(w, templateName, data)
}
