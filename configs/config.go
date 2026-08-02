package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App        AppConfig        `yaml:"app"`
	Server     ServerConfig     `yaml:"server"`
	Pagination PaginationConfig `yaml:"pagination"`
	JWT        JWTConfig        `yaml:"jwt"`
	Database   DatabaseConfig
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Env     string `yaml:"env"`
	Version string `yaml:"version"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type PaginationConfig struct {
	DefaultLimit int `yaml:"defaultLimit"`
}

type DatabaseConfig struct {
	DatabaseUrl string
}

type JWTConfig struct {
	AccessTokenSecret      string
	RefreshTokenSecret     string
	AccessTokenExpiration  time.Duration `yaml:"accessTokenExpiration"`
	RefreshTokenExpiration time.Duration `yaml:"refreshTokenExpiration"`
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("project root containing go.mod not found")
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	rootDir, err := findProjectRoot()
	if err != nil {
		return nil, fmt.Errorf("failed to find project root: %w", err)
	}

	configPath := filepath.Join(rootDir, "configs", "config.yaml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file at %s: %w", configPath, err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config YAML: %w", err)
	}

	// Load environment variables for sensitive data
	cfg.Database.DatabaseUrl = os.Getenv("DATABASE_URL")
	cfg.JWT.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	cfg.JWT.RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")

	if cfg.Database.DatabaseUrl == "" {
		return nil, fmt.Errorf("Database URL is Empty")
	}

	if cfg.JWT.AccessTokenSecret == "" {
		return nil, fmt.Errorf("Access Token Secret is Empty")
	}

	if cfg.JWT.RefreshTokenSecret == "" {
		return nil, fmt.Errorf("Refresh Token Secret is Empty")
	}

	return cfg, nil
}
