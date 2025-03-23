package repository

import "github.com/dongri/crud-go/backend/app/domains/model"

type CommentRepository interface {
	List(id uint64) ([]model.Comment, error)
	New(id uint64, body string) (*model.Comment, error)
}
