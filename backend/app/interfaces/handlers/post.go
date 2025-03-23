package handlers

import (
	"net/http"
	"strconv"

	"github.com/dongri/crud-go/backend/app/interfaces/response"
	"github.com/dongri/crud-go/backend/app/usecases"
	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	usecase usecases.Post
}

func NewPostHandler(u usecases.Post) *PostHandler {
	return &PostHandler{
		usecase: u,
	}
}

type PostJson struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (s PostHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
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

func (s PostHandler) ShowHandler(w http.ResponseWriter, r *http.Request) {
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

func (s PostHandler) NewHandler(w http.ResponseWriter, r *http.Request) {
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

func (s PostHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
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

func (s PostHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
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
