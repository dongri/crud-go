package config

import "os"

type development struct{}

func (c development) PostgresURI() string {
	return os.Getenv("POSTGRES_URI")
}

func (c development) Domain() string {
	return "http://localhost:3001"
}
