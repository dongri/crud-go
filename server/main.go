package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/dongri/crud-go/server/app/infrastructures/postgres"
	"github.com/dongri/crud-go/server/app/interfaces"
	"github.com/dongri/crud-go/server/config"
)

const (
	defaultEnv  = config.Development
	defaultPort = "3001"
)

func main() {
	env := config.Env(os.Getenv("GO_ENV"))
	if env == "" {
		env = config.Env(defaultEnv)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config.SetEnvironment(env)

	// infrastructures
	dbmap := postgres.NewPostgres()
	defer dbmap.Db.Close()

	rooter := interfaces.NewRootHandler(env, dbmap)

	// Launch a server instance.
	fmt.Printf("Server listening on port %s in %s mode.\n", port, env)
	http.ListenAndServe("0.0.0.0:"+port, rooter)
}
