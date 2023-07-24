package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	// API Key for haveibeenpwnd.com
	HibpAPIKey string `mapstructure:"hibpApiKey"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to read .env file: %s", err)
	}
	var config Config
	config.HibpAPIKey = os.Getenv("HIBP_API_KEY")

	return &config, nil
}
