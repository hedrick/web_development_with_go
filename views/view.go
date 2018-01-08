package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// View struct
type View struct {
	Template *template.Template
	Layout   string
}

var (
	// LayoutDir - directory for Globbing
	LayoutDir = "views/layouts/"
	//TemplateExt - extension for Globbing
	TemplateExt = ".gohtml"
)

// NewView to handle logic for each template view
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// Render - renders a view given an interface
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}
