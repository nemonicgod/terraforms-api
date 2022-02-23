package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	databaseVariables = map[string]string{
		"PGUser": "PG_USER",
		"PGPass": "PG_PASS",
		"PGHost": "PG_HOST",
		"PGPort": "PG_PORT",
		"PGDB":   "PG_DB",
	}
)

const (
	// Name for the applicaiton
	Name = "terraforms-api"

	// Env getter for Viper
	Env = "env"

	// Local environment
	Local = "local"

	// Environment getter for Viper
	Environment = "environment"

	// Development environment
	Development = "development"

	// Production environment
	Production = "production"

	// Port variable
	Port = "port"

	// Role variable
	Role = "role"

	PGUser = "POSTGRES_USER"
	PGPass = "POSTGRES_PASS"
	PGHost = "POSTGRES_HOST"
	PGPort = "POSTGRES_PORT"
	PGData = "POSTGRES_DB"

	RedisHost = "REDIS_HOST"
	RedisPort = "REDIS_PORT"

	ROLE_API = "api"
	ROLE_JOB = "job"

	Official = "official"
)

// Reader represents configuration reader
type Reader interface {
	Get(string) interface{}
	GetString(string) string
	GetInt(string) int
	GetBool(string) bool
	GetStringMap(string) map[string]interface{}
	GetStringMapString(string) map[string]string
	GetStringSlice(string) []string
	SetDefault(string, interface{})
}

// DefaultSettings is the function for configuring defaults
type DefaultSettings func(config Reader)

// Defaults is the default settings functor
func Defaults(config Reader) {
	config.SetDefault(Environment, GetEnv("ENVIRONMENT", Development))
	config.SetDefault(Port, GetEnv("PORT", "8000"))
	config.SetDefault(Role, GetEnv("ROLE", "-"))

	config.SetDefault(RedisHost, GetEnv(RedisHost, "127.0.0.1"))
	config.SetDefault(RedisPort, GetEnv(RedisPort, "6379"))

	config.SetDefault(PGUser, GetEnv(PGUser, "postgres"))
	config.SetDefault(PGPass, GetEnv(PGPass, "password"))
	config.SetDefault(PGHost, GetEnv(PGHost, "postgres"))
	config.SetDefault(PGPort, GetEnv(PGPort, "5432"))
	config.SetDefault(PGData, GetEnv(PGData, "postgres"))
}

// GetEnv - pull values or set defaults.
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}

// LoadConfig - returns configuration for a particular app
func LoadConfig(defaultSetup DefaultSettings) (Reader, error) {
	config := viper.New()

	Defaults(config)

	return config, nil
}

// LoadLogger - set the defaults for the logging class
func LoadLoggerConfig(config Reader) *logrus.Logger {
	log := logrus.New()

	log.Formatter = new(prefixed.TextFormatter)
	log.Out = os.Stdout

	log.SetLevel(logrus.DebugLevel)

	return log
}

// LoadLoggerGeneric - returns a generic logger no config required
func LoadLoggerGeneric() *logrus.Logger {
	log := logrus.New()

	log.Formatter = new(prefixed.TextFormatter)
	log.Out = os.Stdout

	log.SetLevel(logrus.InfoLevel)

	return log
}
