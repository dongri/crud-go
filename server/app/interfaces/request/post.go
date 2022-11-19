package request

import (
	"net/http"
	"strconv"

	"github.com/dongri/crud-go/server/app/interfaces/response"
	"github.com/dongri/crud-go/server/app/usecases"
	"github.com/go-chi/chi"
)

type PostRequest struct {
	usecase usecases.Post
}

func NewPostRequest(u usecases.Post) *PostRequest {
	return &PostRequest{
		usecase: u,
	}
}

type PostJson struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (s PostRequest) ListHandler(w http.ResponseWriter, r *http.Request) {
	output := map[string]interface{}{}
	posts, err := s.usecase.List()
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	output["posts"] = posts
	response.JSON(w, r, http.StatusOK, output)
}

func (s PostRequest) ShowHandler(w http.ResponseWriter, r *http.Request) {
	output := map[string]interface{}{}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	post, err := s.usecase.Show(id)
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	output["post"] = post
	response.JSON(w, r, http.StatusOK, output)
}

func (s PostRequest) NewHandler(w http.ResponseWriter, r *http.Request) {
	output := map[string]interface{}{}
	var b PostJson
	if err := readJSON(w, r, &b); err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	post, err := s.usecase.New(b.Title, b.Body)
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	output["post"] = post
	response.JSON(w, r, http.StatusOK, output)
}

func (s PostRequest) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	output := map[string]interface{}{}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	var b PostJson
	if err := readJSON(w, r, &b); err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	post, err := s.usecase.Update(id, b.Title, b.Body)
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	output["post"] = post
	response.JSON(w, r, http.StatusOK, output)
}

func (s PostRequest) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	output := map[string]interface{}{}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	post, err := s.usecase.Delete(id)
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	output["post"] = post
	response.JSON(w, r, http.StatusOK, output)
}
