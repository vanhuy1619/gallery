package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadConfig() {
	pathConfigFile := "config/env.env"
	err := godotenv.Load(pathConfigFile)
	if err != nil {
		log.Print("Load env error")
	}
}
