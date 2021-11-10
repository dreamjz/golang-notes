package config

import "time"

type Server struct {
	HttpPort     int           `mapstructure:"http-port"`
	ReadTimeout  time.Duration `mpstructure:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout"`
}
