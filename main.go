/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"massbit.io/cli/cmd"
)

var Config = ""

func main() {
	if Config != "" {
		viper.SetConfigType("yaml")
		err := viper.ReadConfig(strings.NewReader(Config))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	cmd.Execute()
}
