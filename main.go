package main

import (
	"os"

	"github.com/AlejandroSuero/go-jwt/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
