package config

import "os"

type production struct{}

func (c production) PostgresURI() string {
	return os.Getenv("POSTGRES_URI")
}

func (c production) Domain() string {
	return "https://crud-go.dongri.me"
}
