package cmd

import (
	"github.com/pterm/pterm"
	"github.com/sllt/hole/internal/client/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "hole exec [agent-short-name] [command]",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			pterm.Println("command is required")
			return
		}

		result, err := service.Client.Exec(args[0], args[1])
		if err != nil {
			pterm.Println(err)
			return
		}

		pterm.Println(result)
	},
}
