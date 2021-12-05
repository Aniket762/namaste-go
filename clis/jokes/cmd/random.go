package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get random jokes in your terminal randomly",
	Long:  `Get random jokes in your terminal with various other options randomly`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"json"`
	Status string `json:"status"`
}

func getRandomJoke() {
	fmt.Println("Get random joke 😁")
}

func init() {
	rootCmd.AddCommand(randomCmd)

}
