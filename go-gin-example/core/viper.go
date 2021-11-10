package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-gin-example/global"
	"log"
)

func Viper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file : %w \n", err))
	}
	// watching and re-reading config files
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed, name:%s, op:%s", e.Name, e.Op)
		err = viper.Unmarshal(&global.AppConfig)
		if err != nil {
			log.Printf("Re-reading config failed, %v", err)
		}
		log.Println("New config re-read")
	})
	viper.WatchConfig()
	// decode config into struct
	err = viper.Unmarshal(&global.AppConfig)
	if err != nil {
		log.Fatalf("Unable decode config into struct, %v ", err)
	}
}
