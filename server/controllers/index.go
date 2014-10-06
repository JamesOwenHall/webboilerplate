// Package controllers contains all controller definitions.
package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	. "../common"
)

// Index is the root controller for the server.  It handles displaying the home
// page, serving files from the public directory and managing 404s.
type Index struct{}

func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request, c Context) {
	// Show the home page
	if r.URL.Path == "/" {
		i.showHomepage(w, r, c)
		return
	}

	if i.servePublicFile(w, r, c) {
		return
	}

	NotFound(w, r, c)

	return
}

func (i *Index) showHomepage(w http.ResponseWriter, r *http.Request, c Context) {
	data := NewViewData()
	data.Title = "Hello"
	data.Values["Foo"] = "Button"

	err := c.Execute(w, data, "_header.html", "index.html", "_footer.html")
	if err != nil {
		c.Logger.Errorln(err)
	}
}

func (i *Index) servePublicFile(w http.ResponseWriter, r *http.Request, c Context) bool {
	// Remove the leading slash
	path := filepath.Clean(r.URL.Path)[1:]

	// Convert to OS specific format
	path = filepath.FromSlash(path)

	// Add to public directory
	path = filepath.Join(PublicDir, path)

	// Check if it exists
	if _, err := os.Stat(path); err != nil {
		return false
	}

	// It exists, serve it
	http.ServeFile(w, r, path)

	return true
}
