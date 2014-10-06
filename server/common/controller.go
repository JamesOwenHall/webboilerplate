package common

import (
	"net/http"
)

type Controller interface {
	ServeHTTP(http.ResponseWriter, *http.Request, Context)
}
