package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("ini")
	viper.SetConfigName("config")
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
