package common

import (
	"database/sql"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/sessions"
)

// Context is a struct that is passed with each request to the handling
// controller.  It serves to give the controller access to commonly used
// variables such as a database handle and session information.
type Context struct {
	*Templater
	DB      *sql.DB
	Logger  *logrus.Logger
	Session *sessions.Session
}
