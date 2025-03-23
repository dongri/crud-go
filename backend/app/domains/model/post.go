package model

type Post struct {
	Base
	Title string `db:"title" json:"title"`
	Body  string `db:"body" json:"body"`
}
