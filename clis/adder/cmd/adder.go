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
		takeInputs()
	},
}

func greeter() {
	fmt.Println("Getting into adding two numbers 🙏")
}

func takeInputs() {
	var input1, input2 int
	fmt.Printf("Enter the first number: ")
	fmt.Scan(&input1)
	fmt.Printf("Enter the second number: ")
	fmt.Scan(&input2)
	result := addTwoNumbers(int64(input1), int64(input2))
	fmt.Print("Adding the numbers gives us: ")
	fmt.Print(result)
}

func addTwoNumbers(a int64, b int64) int64 {
	var c int64 = a + b
	return c
}

func init() {
	rootCmd.AddCommand(adderCmd)
}
