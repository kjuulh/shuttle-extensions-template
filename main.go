package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use: "shuttle-extensions-template",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("Hello, shuttle-extensions-template")

			return nil
		},
	}

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
