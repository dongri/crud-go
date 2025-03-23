package handlers

import (
	"net/http"
	"strconv"

	"github.com/dongri/crud-go/backend/app/interfaces/response"
	"github.com/dongri/crud-go/backend/app/usecases"
	"github.com/go-chi/chi/v5"
)

type CommentHandler struct {
	usecase usecases.Comment
}

func NewCommentHandler(u usecases.Comment) *CommentHandler {
	return &CommentHandler{
		usecase: u,
	}
}

type CommentJson struct {
	Body string `json:"body"`
}

func (s CommentHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	output := map[string]interface{}{}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	comments, err := s.usecase.List(id)
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	output["comments"] = comments
	response.JSON(w, r, http.StatusOK, output)
}

func (s CommentHandler) NewHandler(w http.ResponseWriter, r *http.Request) {
	output := map[string]interface{}{}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}

	var b CommentJson
	if err := readJSON(w, r, &b); err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	post, err := s.usecase.New(id, b.Body)
	if err != nil {
		output["error"] = err.Error()
		response.JSON(w, r, http.StatusBadRequest, output)
		return
	}
	output["post"] = post
	response.JSON(w, r, http.StatusOK, output)
}
