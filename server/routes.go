package server

import (
	"./controllers"
)

// setupRoutes defines all routes the server is listening for.  The routes will
// be served by the first pattern that satisfies the request URL.  The root
// controller ("/") will accept all requests, so it should be at the end of the
// list and it should handle non-existing URLs.
func (s *Server) setupRoutes() {
	// Insert routes here

	s.HandleCtrl("/", new(controllers.Index))
}
