package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigType("yaml")
	ymlConfig := []byte(`
app_name: read-config
log_level: debug
mysql:
 port: 3306
redis:
 port: 6379
`)
	err := viper.ReadConfig(bytes.NewBuffer(ymlConfig))
	if err != nil {
		log.Fatal("read config failed:", err)
	}
	fmt.Println("redis port:", viper.GetInt("redis.port"))
}
