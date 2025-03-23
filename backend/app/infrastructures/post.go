package infrastructures

import (
	"fmt"
	"time"

	"github.com/dongri/crud-go/backend/app/domains/model"
	"github.com/dongri/crud-go/backend/app/domains/repository"
	"github.com/go-gorp/gorp/v3"
)

type post struct {
	dbmap *gorp.DbMap
}

func NewPost(dbmap *gorp.DbMap) repository.PostRepository {
	return &post{
		dbmap: dbmap,
	}
}

const (
	postColumns = "id, title, body, created, updated"
)

func (s post) List() ([]model.Post, error) {
	posts := []model.Post{}
	query := fmt.Sprintf(`
		SELECT 
			%s
		FROM posts
		ORDER BY created DESC
	`, postColumns)
	if _, err := s.dbmap.Select(&posts, query); err != nil {
		return nil, err
	}
	return posts, nil
}

func (s post) Get(id uint64) (*model.Post, error) {
	post := new(model.Post)
	query := fmt.Sprintf(`
		SELECT 
			%s
		FROM posts
		WHERE id = $1
	`, postColumns)
	if err := s.dbmap.SelectOne(&post, query, id); err != nil {
		return nil, err
	}
	return post, nil
}

func (s post) New(title, body string) (*model.Post, error) {
	now := time.Now().UTC()
	post := new(model.Post)
	post.Title = title
	post.Body = body
	post.Created = now
	post.Updated = now
	if err := s.dbmap.Insert(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s post) Update(id uint64, title, body string) (*model.Post, error) {
	post, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	post.Title = title
	post.Body = body
	post.Updated = time.Now().UTC()
	if _, err := s.dbmap.Update(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s post) Delete(id uint64) (*model.Post, error) {
	post, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	if _, err := s.dbmap.Delete(post); err != nil {
		return nil, err
	}
	return post, nil
}
