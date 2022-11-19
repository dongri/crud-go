package usecases

import (
	"github.com/dongri/crud-go/server/app/domains/model"
	"github.com/dongri/crud-go/server/app/domains/repository"
)

type Post struct {
	repository repository.PostRepository
}

func NewPost(repository repository.PostRepository) Post {
	return Post{
		repository: repository,
	}
}

func (s Post) List() ([]model.Post, error) {
	posts, err := s.repository.List()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s Post) Show(id uint64) (*model.Post, error) {
	post, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s Post) New(title, body string) (*model.Post, error) {
	post, err := s.repository.New(title, body)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s Post) Update(id uint64, title, body string) (*model.Post, error) {
	post, err := s.repository.Update(id, title, body)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s Post) Delete(id uint64) (*model.Post, error) {
	post, err := s.repository.Delete(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}
