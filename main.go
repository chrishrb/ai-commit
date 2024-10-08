/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log/slog"

	"github.com/chrishrb/ai-commit/cmd"
)

func main() {
  slog.SetLogLoggerLevel(slog.LevelDebug)
	cmd.Execute()
}
