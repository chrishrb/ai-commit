package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chrishrb/ai-commit/pkg/client"
	"github.com/spf13/cobra"
)

// hookCmd represents the hook command
var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Commit hook",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("hook not executed correct")
			return
		}

		var commitMsgFile = args[0]
		response, err := client.BuildCommitMessage()
		if err != nil {
			log.Fatal(err)
			return
		}

		file, err := os.OpenFile(commitMsgFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModeAppend)
		if _, err := file.WriteString(response); err != nil {
			log.Fatal(err)
			return
		}
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
