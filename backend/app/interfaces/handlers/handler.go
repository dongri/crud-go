package handlers

import (
	"github.com/dongri/crud-go/backend/app/usecases"
)

type Hnadler struct {
	PostHandler    *PostHandler
	CommentHandler *CommentHandler
}

func NewHandler(u usecases.Usecase) *Hnadler {
	return &Hnadler{
		PostHandler:    NewPostHandler(u.PostUseCase),
		CommentHandler: NewCommentHandler(u.CommentUseCase),
	}
}
