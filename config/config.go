package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	configName := flag.String("config", "config", "config file name")
	flag.Parse()

	viper.SetConfigType("ini")
	viper.SetConfigName(*configName)
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetString(s string) string {
	return viper.GetString(s)
}
