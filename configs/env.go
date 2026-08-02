package configs

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	rootDir, err := findProjectRoot()
	if err != nil {
		log.Printf("Warning: could not find project root: %v. Falling back to default .env loading.", err)
		err = godotenv.Load()
	} else {
		err = godotenv.Load(filepath.Join(rootDir, ".env"))
	}
	if err != nil {
		log.Fatal("Error loading env: ", err)
	}
}
