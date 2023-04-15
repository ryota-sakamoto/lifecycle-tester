package main

import (
	"github.com/ryota-sakamoto/lifecycle-tester/internal/cmd"
)

func main() {
	command := cmd.NewCommand()
	if err := command.Execute(); err != nil {
		panic(err)
	}
}
