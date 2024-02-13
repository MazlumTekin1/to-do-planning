package config

import "os"

type Config struct {
	Provider1URL string
	Provider2URL string
}

func LoadConfig() *Config {

	provider1URL := os.Getenv("PROVIDER1_URL")
	if provider1URL == "" {
		provider1URL = "https://run.mocky.io/v3/27b47d79-f382-4dee-b4fe-a0976ceda9cd"
	}

	provider2URL := os.Getenv("PROVIDER2_URL")
	if provider2URL == "" {
		provider2URL = "https://run.mocky.io/v3/7b0ff222-7a9c-4c54-9396-0df58e289143"
	}

	return &Config{
		Provider1URL: provider1URL,
		Provider2URL: provider2URL,
	}
}
