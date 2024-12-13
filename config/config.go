package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mekonger/go-image-generator/global"
	"os"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		global.ServerMode = "DEV"
		fmt.Println("Error loading .env file")
	}

	global.ServerMode = os.Getenv("SERVER_MODE")
}
