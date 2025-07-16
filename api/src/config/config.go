package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnection = ""
	Port             = 0
)

func Load() {
	var error error

	if error = godotenv.Load(); error != nil {
		panic("Error loading .env file")
	}

	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		Port = 9000
	}

	StringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

}
