package cmd

import (
	"github.com/spf13/cobra"
	"go-api-template/cmd/commands"
)

var (
	rootCmd = &cobra.Command{
		Use:   "country-info",
		Short: "A geography game",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(commands.CapitalsCmd)
}

func initConfig() {
}