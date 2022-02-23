package backend

import (

	// "strconv"

	"github.com/go-redis/redis/v8"
	"github.com/nemonicgod/terraforms-api/config"
	"gorm.io/gorm"

	db "github.com/nemonicgod/terraforms-api/infra/database"
	rd "github.com/nemonicgod/terraforms-api/infra/redis"

	"github.com/sirupsen/logrus"
)

// Repository ...
type Repository struct {
	D *gorm.DB
	R *redis.Client
}

// Backend - main struct for the entire application configuration
type Backend struct {
	// C - contains the yaml file configuration key/values and other env specifics
	C config.Reader

	// L - a logrus logger, customized for this application
	L *logrus.Logger

	// R - a repository object for holding db/redis connections
	R *Repository
}

// NewBackend - factory method for producing a new type of Backend
func NewBackend(cDB bool, cRD bool) (*Backend, error) {
	c, err := config.LoadConfig(config.Defaults)
	if err != nil {
		return nil, err
	}

	var dss *gorm.DB
	var rdd *redis.Client

	if cDB {
		// Database connection
		dss = db.Connect(c)
	}

	if cRD {
		// Redis connection
		rdd = rd.Connect(c)
	}

	// Base BackendConfiguration to link structs and objects
	var bc = &Backend{
		C: c,
		L: config.LoadLoggerConfig(c),
		R: &Repository{
			D: dss,
			R: rdd,
		},
	}

	return bc, nil
}
