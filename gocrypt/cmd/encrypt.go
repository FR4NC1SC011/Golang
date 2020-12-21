/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a file",
	Run: func(cmd *cobra.Command, args []string) {
		secret, err := Secret()
		if err != nil {
			log.Fatal(err)
		}
		file := args[0]
		message, err := Read_file(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Message:", message)

		encrypted, err := Encrypt(message, secret)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Encrypted:", encrypted)

	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
