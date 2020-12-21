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
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a file",
	Run: func(cmd *cobra.Command, args []string) {

		file := args[0]                 // File to encrypt is the first arg.
		message, err := Read_file(file) // Read File Function on data_encryption.go
		Check(err)                      // Function on data_encryption.go

		secret, err := ioutil.ReadFile(args[1]) // Read second argument as a string (Secret)
		if err != nil {
			fmt.Println("key.key not specified")
			log.Fatal(err)
		}

		encrypted, err := Encrypt(message, secret) // Encrypt File Using the secret (2 arg)
		Check(err)

		enc_msg := []byte(encrypted)
		err = ioutil.WriteFile("file_encrypted.txt", enc_msg, 0644)
		Check(err)

		fmt.Println("Encrypted:", encrypted)

	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
