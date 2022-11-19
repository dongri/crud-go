package config

type development struct{}

func (c development) PostgresURI() string {
	return "postgresql://postgres@crud-go_postgres:5432/crud-go?sslmode=disable"
}

func (c development) Domain() string {
	return "http://localhost:3001"
}
