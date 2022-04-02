package main

import (
	"github.com/dewadg/go-playground-api/cmd/generateids"
	"github.com/dewadg/go-playground-api/cmd/serve"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{
		Use: "go-playground",
	}

	rootCmd.AddCommand(serve.Command())
	rootCmd.AddCommand(generateids.Command())

	_ = rootCmd.Execute()
}
