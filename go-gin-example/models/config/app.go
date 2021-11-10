package config

type App struct {
	PageSize  int    `mapstructure:"page-size"`
	JwtSecret string `mapstructure:"jwt-secret"`
}
