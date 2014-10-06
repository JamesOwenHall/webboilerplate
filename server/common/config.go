// Package common holds all code that should be accessible from any other
// package.  It has no dependencies within the server package.
package common

// Site directories.
const (
	PublicDir = "public"
	ViewsDir  = "views"
)

// Session information.
const (
	SessionName   = "default"
	SessionSecret = "my-session-secret"
)
