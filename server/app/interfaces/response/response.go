package response

import (
	"net/http"

	renderLib "github.com/unrolled/render"
)

var render = renderLib.New()

func JSON(w http.ResponseWriter, r *http.Request, httpStatus int, data interface{}) {
	if err := render.JSON(w, httpStatus, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
