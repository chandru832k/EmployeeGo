package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type environmentData struct {
	DbUrl  string `validate:"required"`
	DbUser string `validate:"required"`
	DbPass string `validate:"required"`
	DbPort string `validate:"required"`
	DbName string `validate:"required"`
}

var EnvironmentData *environmentData

func InitializeEnv() {
	_ = godotenv.Load("./Config.env")
	EnvironmentData = &environmentData{
		DbUrl:  os.Getenv("DB_URL"),
		DbName: os.Getenv("DB_NAME"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
		DbPort: os.Getenv("DB_PORT"),
	}
	fmt.Println("%v", EnvironmentData)
}
