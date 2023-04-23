/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lw := list.NewWriter()
		for _, blockchain := range config.Blockchains {
			lw.AppendItem(blockchain.Name)
			lw.Indent()
			for _, network := range blockchain.Networks {
				lw.AppendItem(network)
			}
			lw.UnIndent()
		}
		fmt.Printf("%s\n", lw.Render())
	},
}

func init() {
	blockchainCmd.AddCommand(listCmd)
}
