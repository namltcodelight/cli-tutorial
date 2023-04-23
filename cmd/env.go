/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var config *Config

type Config struct {
	AppEnv      string
	Blockchains map[string]Blockchain
}

type Blockchain struct {
	Name     string
	Networks []string
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.AppEnv)
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
