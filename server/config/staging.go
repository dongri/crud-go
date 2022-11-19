package config

import "os"

type staging struct{}

func (c staging) PostgresURI() string {
	return os.Getenv("POSTGRES_URI")
}

func (c staging) Domain() string {
	return "http://localhost:3001"
}
