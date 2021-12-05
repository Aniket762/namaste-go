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

	"github.com/spf13/cobra"
)

// adderCmd represents the adder command
var adderCmd = &cobra.Command{
	Use:   "adder",
	Short: "A simple cli to add two numbers",
	Long:  `A simple cli to add two numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		greeter()

	},
}

func greeter() {
	fmt.Println("Getting into adding two numbers 🙏")
}

func init() {
	rootCmd.AddCommand(adderCmd)
}
