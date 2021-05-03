package apiserver

//Config is server settings
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DataBaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

//NewConfig returns default config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
