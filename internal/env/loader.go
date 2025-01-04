package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lakpahana/adaptive-video-uploader/internal/path"
)

// LoadEnv loads environment variables from a `.env` file
func LoadEnv() {
	rootPath := path.GetProjectRootPath()
	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(-1)
	}
}
