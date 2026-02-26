package config

import "os"

type Config struct {
	AppName   string
	Port      string
	JWTSecret string
}

func Load() Config {
	return Config{
		AppName:   getEnv("APP_NAME", "NusaStay"),
		Port:      getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "dev-secret"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
