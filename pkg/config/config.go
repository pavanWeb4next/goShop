package config

import (
	"log"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

const (
	ProductionEnv = "production"

	DatabaseTimeout    = 5 * time.Second
	ProductCachingTime = 1 * time.Minute
)

var AuthIgnoreMethods = []string{
	"/user.UserService/Login",
	"/user.UserService/Register",
}

type Schema struct {
	Environment   string `env:"environment"`
	HttpPort      int    `env:"http_port"`
	GrpcPort      int    `env:"grpc_port"`
	AuthSecret    string `env:"auth_secret"`
	DatabaseURI   string `env:"database_uri"`
	RedisURI      string `env:"redis_uri"`
	RedisPassword string `env:"redis_password"`
	RedisDB       int    `env:"redis_db"`
}

var (
	cfg Schema
)

func LoadConfig() *Schema {
	// Load .env from root project directory
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file, error: %v", err)
	}

	// Parse environment variables into cfg struct
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error parsing environment variables: %v", err)
	}

	return &cfg
}

func GetConfig() *Schema {
	return &cfg
}
