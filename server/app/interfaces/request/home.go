package request

import (
	"net/http"

	"github.com/dongri/crud-go/server/app/interfaces/response"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	output := map[string]interface{}{}
	output["api"] = "crud-go"
	response.JSON(w, r, http.StatusOK, output)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	output := map[string]interface{}{}
	output["version"] = "1.0.0"
	response.JSON(w, r, http.StatusOK, output)
}
