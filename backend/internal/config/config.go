package config

type Config struct {
	ServerPort string
	AppName    string
}

func LoadConfig() Config {
	return Config{
		ServerPort: "8080",
		AppName:    "DCP Backend",
	}
}
