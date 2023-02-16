package config

type Config struct {
	KeysDir string
}

func NewConfig(keysDir string) *Config {
	return &Config{
		KeysDir: keysDir,
	}
}
