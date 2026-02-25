package config

type Config struct {
	AppName string
	Port    string
}

func Load() Config {
	return Config{AppName: "NusaStay", Port: GetEnv("APP_PORT", "8080")}
}
