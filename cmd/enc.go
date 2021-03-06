/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/nevilleomangi/vigenere-cypher/vigenere"
	"github.com/spf13/cobra"
)

// encCmd represents the enc command
var (
	encCmd = &cobra.Command{
		Use:   "enc",
		Short: "Encrypt a message using a key",
		Long: `Usage:

	vinegere enc -k <key> <message>
	`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", vigenere.Encrypt(keyFlagValue, args[0]))
		},
	}
	keyFlagValue string
)

func init() {
	rootCmd.AddCommand(encCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	encCmd.Flags().StringVarP(&keyFlagValue, "key", "k", "", "Key to encrypt the message")
	encCmd.MarkFlagRequired("key")
}
