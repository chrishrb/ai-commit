/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/chrishrb/ai-commit/pkg/client"
	"github.com/spf13/cobra"
)

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
    response, err := client.GetResponse()
    if err != nil {
      log.Fatal(err)
      return
    }
    fmt.Println(*response)
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
