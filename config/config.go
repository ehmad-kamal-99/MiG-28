package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config container for all config.
type Config struct {
	Deploy         *string `envconfig:"DEPLOY" default:"dev"`
	ProjectID      *string `envconfig:"PROJECT_ID" required:"true"`
	CommitSHA      *string `envconfig:"COMMIT_SHA" required:"true"`
	LogLevel       *string `envconfig:"LOG_LEVEL" default:"1"`
	ServiceConfig  *ServiceConfig
	FirebaseConfig FirebaseConfig
	MongoDBConfig  *MongoDBConfig
	RollbarConfig  *RollbarConfig
	FastlyConfig   *FastlyConfig
}

// ServiceConfig contains configs related to current service.
type ServiceConfig struct {
	ServiceName     *string `envconfig:"K_SERVICE" required:"true"`
	ServiceRevision *string `envconfig:"K_REVISION" required:"true"`
}

// MongoDBConfig holds config for mongo database.
type MongoDBConfig struct {
	MongoDBURI *string `envconfig:"MONGO_DB_URI" required:"true"`
}

// FirebaseConfig contains configs for firebase auth.
type FirebaseConfig struct {
	ServiceAccountID *string `envconfig:"SERVICE_ACCOUNT_ID" required:"true"`
}

// RollbarConfig contains configs for rollbar error reporting.
type RollbarConfig struct {
	RollbarToken      *string `envconfig:"ROLLBAR_TOKEN" required:"true"`
	RollbarEnv        *string `envconfig:"ROLLBAR_ENV" default:"development"`
	RollbarServerRoot *string `envconfig:"ROLLBAR_SERVER_ROOT" default:"/go/src/xyz/code/user/"`
}

// FastlyConfig contains configs for fastly caching service.
type FastlyConfig struct {
	FastlyServiceName string `envconfig:"FASTLY_SERVICE_NAME" required:"true"`
	FastlyToken       string `envconfig:"FASTLY_TOKEN" required:"true"`
}

// LoadConfig reads the environment variables and populates the object.
func LoadConfig() Config {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
