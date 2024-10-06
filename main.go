/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/chrishrb/ai-commit/cmd"
)

func main() {
  log.SetPrefix("ai-commit: ")
	cmd.Execute()
}
