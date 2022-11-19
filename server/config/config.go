package config

import (
	"fmt"
)

// Env ...
type Env string

// Const ...
const (
	Development Env = "development"
	Staging     Env = "staging"
	Production  Env = "production"
)

// CurrentEnv ...
var CurrentEnv Env

func (e Env) String() string {
	return string(e)
}

// EnvConfig ...
type EnvConfig interface {
	PostgresURI() string

	Domain() string
}

// CurrentConfig ...
var CurrentConfig EnvConfig

// SetEnvironment ..
func SetEnvironment(e Env) error {
	var configInstance EnvConfig
	switch e {
	case Production:
		configInstance = &production{}
	case Staging:
		configInstance = &staging{}
	case Development:
		configInstance = &development{}
	}
	if configInstance == nil {
		return fmt.Errorf("no configuration found for %s", e)
	}

	CurrentConfig = configInstance
	CurrentEnv = e
	return nil
}

func IsLocal() bool {
	return CurrentEnv == Development
}

func IsProduction() bool {
	return CurrentEnv == Production
}
