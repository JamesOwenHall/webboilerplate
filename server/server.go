// Package server is the root package of all Go server code.
package server

import (
	"database/sql"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/sessions"
	"net/http"

	. "./common"
)

// Server contains all necessary information to host an instance of the
// website.  You will typically only create one instance of this struct.
type Server struct {
	http.Server
	db        *sql.DB
	logger    *logrus.Logger
	mux       *http.ServeMux
	store     sessions.Store
	templater Templater
}

// New initializes and returns a new instance of a Server.
func New(port int) *Server {
	result := new(Server)
	result.Addr = fmt.Sprint(":", port)

	// Setup logger
	result.logger = logrus.New()
	result.logger.Level = logrus.DebugLevel

	// Setup database
	var err error
	result.db, err = setupDatabase()
	if err != nil {
		result.logger.Panicln(err)
	}

	// Setup templater
	result.templater = Templater{
		Caching:   false,
		Directory: ViewsDir,
	}
	result.templater.Reset()

	// Setup session
	result.store = sessions.NewCookieStore([]byte(SessionSecret))

	// Setup mux
	result.mux = http.NewServeMux()
	result.setupRoutes()
	result.Handler = result.mux

	return result
}

// ListenAndServe calls the underlying http.Server's ListenAndServe function,
// with the added benefit of logging this action.
func (s *Server) ListenAndServe() error {
	s.logger.Infoln("Listening on", s.Addr)

	if s.templater.Caching {
		s.logger.Infoln("Template caching is enabled")
	} else {
		s.logger.Infoln("Template caching is disabled")
	}

	return s.Server.ListenAndServe()
}

// HandleCtrl adds a route to the specified controller, similar to http.Handle.
func (s *Server) HandleCtrl(pattern string, ctrl Controller) {
	s.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.store.Get(r, SessionName)

		c := Context{
			Logger:    s.logger,
			Session:   session,
			Templater: &s.templater,
		}

		ctrl.ServeHTTP(w, r, c)
	})
}

// SetTemplateCaching allows you to enable or disable template caching.  By
// default, caching is disabled.  SetTemplateCaching should only be called
// before ListenAndServe is called.
func (s *Server) SetTemplateCaching(enabled bool) {
	s.templater.Caching = enabled
	s.templater.Reset()
}
