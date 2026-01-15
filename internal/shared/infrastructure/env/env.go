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
	R2        R2Config
}

type R2Config struct {
	AccessKeyID     string
	SecretAccessKey string
	AccountID       string
	BucketName      string
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

	r2Config := R2Config{
		AccessKeyID:     os.Getenv("R2_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("R2_SECRET_ACCESS_KEY"),
		AccountID:       os.Getenv("R2_ACCOUNT_ID"),
		BucketName:      os.Getenv("R2_BUCKET_NAME"),
	}

	return &Config{
		MongoURI:  uri,
		ServerEnv: serverEnv,
		R2:        r2Config,
	}, nil

}
