package usecases

import (
	"github.com/dongri/crud-go/backend/app/domains/model"
	"github.com/dongri/crud-go/backend/app/infrastructures"
)

type Post struct {
	infra infrastructures.Infra
}

func NewPost(infra infrastructures.Infra) Post {
	return Post{
		infra: infra,
	}
}

func (s Post) List() ([]model.Post, error) {
	posts, err := s.infra.PostRepository.List()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s Post) Show(id uint64) (*model.Post, error) {
	post, err := s.infra.PostRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s Post) New(title, body string) (*model.Post, error) {
	post, err := s.infra.PostRepository.New(title, body)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s Post) Update(id uint64, title, body string) (*model.Post, error) {
	post, err := s.infra.PostRepository.Update(id, title, body)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s Post) Delete(id uint64) (*model.Post, error) {
	post, err := s.infra.PostRepository.Delete(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}
