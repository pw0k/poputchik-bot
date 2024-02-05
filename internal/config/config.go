package config

import "os"

type Config struct {
	BotToken string
}

func LoadConfig() *Config {
	return &Config{
		BotToken: os.Getenv("STRANGER_TELEGRAMBOT_TOKEN"),
	}
}
