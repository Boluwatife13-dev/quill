package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	MongoURI  string
	JWTSecret string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return Config{
		Port:      os.Getenv("PORT"),
		MongoURI:  os.Getenv("MONGO_URI"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}