package infrastructures

import (
	"fmt"
	"time"

	"github.com/dongri/crud-go/backend/app/domains/model"
	"github.com/dongri/crud-go/backend/app/domains/repository"
	"github.com/go-gorp/gorp/v3"
)

type comment struct {
	dbmap *gorp.DbMap
}

func NewComment(dbmap *gorp.DbMap) repository.CommentRepository {
	return &comment{
		dbmap: dbmap,
	}
}

const (
	commentColumns = "id, post_id, body, created, updated"
)

func (s comment) List(id uint64) ([]model.Comment, error) {
	comments := []model.Comment{}

	query := fmt.Sprintf(`
		SELECT
			%s
		FROM comments
		WHERE post_id = $1
		ORDER BY created DESC
	`, commentColumns)

	if _, err := s.dbmap.Select(&comments, query, id); err != nil {
		return nil, err
	}
	return comments, nil
}

func (s comment) New(id uint64, body string) (*model.Comment, error) {
	now := time.Now().UTC()
	comment := new(model.Comment)
	comment.PostID = id
	comment.Body = body
	comment.Created = now
	comment.Updated = now
	if err := s.dbmap.Insert(comment); err != nil {
		return nil, err
	}
	return nil, nil
}
