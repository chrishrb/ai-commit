/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chrishrb/ai-commit/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:           "ai-commit",
		Short:         "ai-commit – command-line tool to generate commit messages using ai",
		Long:          ``,
		Version:       "0.1.0",
		SilenceErrors: true,
		SilenceUsage:  true,
	  Run: func(cmd *cobra.Command, args []string) {
      _, err := client.BuildCommitMessage()
      if err != nil {
        log.Fatal(err)
        return
      }
    },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
  cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ai-commit.yaml)")
  viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

func initConfig() {
  if cfgFile != "" {
    viper.SetConfigFile(cfgFile)
  } else {
    home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigFile(".ai-commit")
  }

  viper.AutomaticEnv()

  if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config file used for jamctl: ", viper.ConfigFileUsed())
	}
}


