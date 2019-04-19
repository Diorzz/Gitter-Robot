package config

import "github.com/Unknwon/goconfig"

type Config struct {
	GitterToken string
}

func LoadConfig() *Config {
	config := &Config{}
	cfg, err := goconfig.LoadConfigFile("config.cfg")
	if err != nil {
		panic("Can't load config file")
	}

	// Load gitter.token
	token, err := cfg.GetValue("gitter", "token")
	if err != nil {
		panic("Cant't load config: [gitter] -> token")
	}
	config.GitterToken = token

	return config
}
