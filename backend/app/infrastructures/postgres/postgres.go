package postgres

import (
	"database/sql"
	"log"

	"github.com/dongri/crud-go/backend/app/domains/model"
	"github.com/dongri/crud-go/backend/config"
	"github.com/go-gorp/gorp/v3"
)

func NewPostgres() *gorp.DbMap {
	db, err := sql.Open("postgres", config.CurrentConfig.PostgresURI())
	if err != nil {
		log.Fatalf("Failed to connect postgres %v", err)
		return nil
	}
	// defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping postgres %v", err)
		return nil
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	addTableWithName(dbmap)

	return dbmap
}

func addTableWithName(dbmap *gorp.DbMap) {
	dbmap.AddTableWithName(model.Post{}, "posts").SetKeys(true, "ID")
	dbmap.AddTableWithName(model.Comment{}, "comments").SetKeys(true, "ID")
}
