package main

import (
	"fmt"
	viper2 "github.com/spf13/viper"
)

func main() {
	viper2.SetConfigFile("config.yml")
	err := viper2.ReadInConfig()
	if err != nil {
		return
	}

	fmt.Print(viper2.GetStringSlice("whitelist"))
}
