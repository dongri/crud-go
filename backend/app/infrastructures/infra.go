package infrastructures

import (
	"github.com/dongri/crud-go/backend/app/domains/repository"
	"github.com/go-gorp/gorp/v3"
)

type Infra struct {
	PostRepository    repository.PostRepository
	CommentRepository repository.CommentRepository
}

func NewInfra(
	dbmap *gorp.DbMap,
) Infra {
	return Infra{
		PostRepository:    NewPost(dbmap),
		CommentRepository: NewComment(dbmap),
	}
}
