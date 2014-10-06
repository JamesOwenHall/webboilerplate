package common

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// Templater allows you to easily enable or disable template caching.  Note,
// it does not support nested templates.  All templates must end in .html
type Templater struct {
	Caching   bool
	Directory string
	tmpl      *template.Template
}

// Reset must be called for new settings to take effect.  It should be called
// once the Templater is created.
func (t *Templater) Reset() error {
	var err error
	t.tmpl = nil

	// Cache if necessary
	if t.Caching {
		t.tmpl, err = template.ParseGlob(filepath.Join(t.Directory, "*.html"))
	}

	return err
}

// Execute is similar to the Execute method of *template.Template.  However,
// this accepts multiple templates which will be executed in sequence and it
// supports caching.
func (t *Templater) Execute(w http.ResponseWriter, data ViewData, names ...string) error {
	if t.Caching {
		for _, n := range names {
			err := t.tmpl.ExecuteTemplate(w, n, data)
			if err != nil {
				return err
			}
		}
	} else {
		for _, n := range names {
			fullName := filepath.Join(t.Directory, n)

			temp, err := template.ParseFiles(fullName)
			if err != nil {
				return err
			}

			err = temp.Execute(w, data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
