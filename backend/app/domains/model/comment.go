package model

type Comment struct {
	Base
	PostID uint64 `db:"post_id" json:"post_id"`
	Body   string `db:"body" json:"body"`
}
