package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var Config Reader

func init() {
	os.Setenv("ENVIRONMENT", "test")
	Config, _ = LoadConfig(Defaults)
	os.Setenv("ENVIRONMENT", "")

	if Config.GetString("environment") != "test" {
		panic(fmt.Errorf("test [environment] is not in [test] mode"))
	}
}

func TestConstants(t *testing.T) {
	assert.Equal(t, Env, "env")
	assert.Equal(t, Environment, "environment")

	assert.Equal(t, Production, "production")
	assert.Equal(t, Development, "development")

	assert.Equal(t, Port, "port")
}

func TestLoadLogger(t *testing.T) {
	logger := LoadLoggerConfig(Config)

	if assert.NotNil(t, logger) {
		assert.Equal(t, logger.Formatter, &prefixed.TextFormatter{})
		assert.Equal(t, logger.Level, logrus.DebugLevel)
	}
}

func TestGetEnvExists(t *testing.T) {
	os.Setenv("FOO", "nothing")

	assert.Equal(t, GetEnv("FOO", "invalid"), "nothing")

	os.Unsetenv("FOO")
}

func TestGetEnvNotExists(t *testing.T) {
	os.Setenv("FOO", "")

	assert.Equal(t, GetEnv("FOO", "invalid"), "invalid")

	os.Unsetenv("FOO")
}
