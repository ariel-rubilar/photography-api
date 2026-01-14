package env

import (
	"fmt"
	"os"

	"github.com/ariel-rubilar/photography-api/internal/server"
	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI  string
	ServerEnv server.Env
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

	var serverEnv server.Env = server.Production

	if env != "production" {
		serverEnv = server.Development
	}

	return &Config{
		MongoURI:  uri,
		ServerEnv: serverEnv,
	}, nil

}
