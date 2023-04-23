/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		Ask("Username", func(s string) error {
			if matched, _ := regexp.MatchString("^[a-zA-Z0-9]{3,10}$", s); !matched {
				return errors.New("Username must be at least 3 characters (max 10 character) includes lowercase or capitals or number")
			}
			return nil
		})
		SelectInList("Gender", []string{"Male", "Female"}, 0)
		password, err := AskSecret("Password", func(s string) error {
			if matched, _ := regexp.MatchString("^[a-zA-Z0-9]{6,10}$", s); !matched {
				return errors.New("Password must be at least 6 characters (max 10 character) includes lowercase or capitals or number")
			}
			return nil
		})
		if err != nil {
			cobra.CheckErr(err)
		}
		SecretConfirm("Confirm Password", password)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}

func Ask(question string, validate promptui.ValidateFunc) (answer string, err error) {
	prompt := promptui.Prompt{
		Label:    question,
		Validate: validate,
	}
	answer, err = prompt.Run()
	return
}

func AskSecret(question string, validate promptui.ValidateFunc) (answer string, err error) {
	prompt := promptui.Prompt{
		Label:    question,
		Mask:     '*',
		Validate: validate,
	}
	answer, err = prompt.Run()
	return
}

func SecretConfirm(question string, confirmation string) (answer string, err error) {
	prompt := promptui.Prompt{
		Label: question,
		Validate: func(input string) error {
			if input != confirmation {
				return errors.New("wrong password")
			}
			return nil
		},
		Mask: '*',
	}
	answer, err = prompt.Run()
	return
}

func SelectInList(text string, selects []string, cursorPos int) (int, error) {
	prompt := promptui.Select{
		Label: text,
		Items: selects,
		Size:  10,
		Searcher: func(input string, index int) bool {
			return strings.Contains(strings.ToLower(selects[index]), strings.ToLower(input))
		},
		CursorPos: 0,
	}
	i, _, err := prompt.Run()
	return i, err
}
