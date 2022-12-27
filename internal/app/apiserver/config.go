package apiserver

type Config struct {
	BindAddr string `toml: bind_addr`
}

//New config
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
	}
}
