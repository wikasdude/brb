package config

type Config struct {
	DBURL string
}

func Load() *Config {
	return &Config{
		DBURL: "postgres://user:password@localhost:5432/brb",
	}
}
