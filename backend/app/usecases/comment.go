package usecases

import (
	"github.com/dongri/crud-go/backend/app/domains/model"
	"github.com/dongri/crud-go/backend/app/infrastructures"
)

type Comment struct {
	infra infrastructures.Infra
}

func NewComment(infra infrastructures.Infra) Comment {
	return Comment{
		infra: infra,
	}
}

func (s Comment) List(id uint64) ([]model.Comment, error) {
	posts, err := s.infra.CommentRepository.List(id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s Comment) New(id uint64, body string) (*model.Comment, error) {
	post, err := s.infra.CommentRepository.New(id, body)
	if err != nil {
		return nil, err
	}
	return post, nil
}
