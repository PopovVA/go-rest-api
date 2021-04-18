package store

//Config is
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

//NewConfig is ..
func NewConfig() *Config {
	return &Config{}
}
