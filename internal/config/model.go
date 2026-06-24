package config

type Config struct {
	Database Database `toml:"database"`
	Server   Server   `toml:"server"`
}

type Database struct {
	Name string `toml:"name"`
}

type Server struct {
	Port            int `toml:"port"`
	ShutdownTimeout int `toml:"shutdown_timeout"`
}
