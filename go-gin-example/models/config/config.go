package config

type Config struct {
	RunMode  string   `mapstructure:"run-mode"`
	App      App      `mapstructure:"app"`
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}
