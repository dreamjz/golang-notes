package config

type Server struct {
	HttpPort     string `mapstructure:"http-port"`
	ReadTimeout  int    `mpstructure:"read-timeout"`
	WriteTimeout int    `mapstructure:"write-timeout"`
}
