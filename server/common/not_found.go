package common

import (
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request, c Context) {
	w.WriteHeader(http.StatusNotFound)

	data := NewViewData()
	data.Title = "Page Not Found"

	err := c.Execute(w, data, "_header.html", "notfound.html", "_footer.html")
	if err != nil {
		c.Logger.Errorln(err)
	}
}
