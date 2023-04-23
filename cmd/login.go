/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	email, password string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err error
		)
		reader := bufio.NewReader(os.Stdin)
		if email == "" {
			fmt.Print("Email: ")
			email, err = reader.ReadString('\n')
			if err != nil {
				cobra.CheckErr(err)
				return
			}
		}

		if password == "" {
			fmt.Print("Password: ")
			passwordB, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				cobra.CheckErr(err)
				return
			}
			password = string(passwordB)
		}

		_, err = Login(strings.TrimSpace(email), string(password))
		if err != nil {
			cobra.CheckErr(err)
		} else {
			fmt.Println("\nLogin success")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&email, "email", "e", "", "Email")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
}

func Login(email string, password string) (accessToken string, err error) {
	values := map[string]string{"email": email, "password": password}
	jsonData, err := json.Marshal(values)

	if err != nil {
		return
	}
	// Get the data
	response, err := http.Post("https://staging.portal.massbitroute.net/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	defer response.Body.Close()
	var res map[string]interface{}

	if response.StatusCode == 200 || response.StatusCode == 201 {
		err = json.NewDecoder(response.Body).Decode(&res)
		if err != nil {
			return
		}
		accessToken = res["accessToken"].(string)
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("error with status %v %v", response.StatusCode, string(body))
	}
	return
}
