package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
}

func LoadConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		return nil, fmt.Errorf("MONGO_URI environment variable is not set")
	}

	return &Config{
		MongoURI: uri,
	}, nil

}
