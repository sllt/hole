package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"hole/internal/client/service"
)

func init() {
	rootCmd.AddCommand(exitCmd)
}

var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "hole exit [agent-short-name]",
	Run: func(cmd *cobra.Command, args []string) {

		result, _ := service.Client.Exit(args[0])

		pterm.Println(result)
	},
}
