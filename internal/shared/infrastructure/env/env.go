package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env string

const (
	Development Env = "development"
	Production  Env = "production"
)

type Config struct {
	MongoURI  string
	ServerEnv Env
}

func LoadConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		return nil, fmt.Errorf("MONGO_URI environment variable is not set")
	}

	env := os.Getenv("SERVER_ENV")

	var serverEnv Env = Production

	if env != "production" {
		serverEnv = Development
	}

	return &Config{
		MongoURI:  uri,
		ServerEnv: serverEnv,
	}, nil

}
