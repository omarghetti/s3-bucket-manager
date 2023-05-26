package setting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	aws_access_key     string
	aws_secret_key     string
	aws_default_region string
}

var Settings = &Env{}

func SetupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env vars from .env file")
	}
	Settings.aws_access_key = os.Getenv("AWS_ACCESS_KEY_ID")
	Settings.aws_secret_key = os.Getenv("AWS_SECRET_KEY_ID")
	Settings.aws_default_region = os.Getenv("AWS_DEFAULT_REGION")
}
