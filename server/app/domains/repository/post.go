package repository

import "github.com/dongri/crud-go/server/app/domains/model"

type PostRepository interface {
	List() ([]model.Post, error)
	Get(id uint64) (*model.Post, error)
	New(title, body string) (*model.Post, error)
	Update(id uint64, title, body string) (*model.Post, error)
	Delete(id uint64) (*model.Post, error)
}
