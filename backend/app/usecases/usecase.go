package usecases

import (
	"github.com/dongri/crud-go/backend/app/infrastructures"
)

type Usecase struct {
	PostUseCase    Post
	CommentUseCase Comment
}

func NewUseeCase(
	infra infrastructures.Infra,
) Usecase {
	return Usecase{
		PostUseCase:    NewPost(infra),
		CommentUseCase: NewComment(infra),
	}
}
