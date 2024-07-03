package cmd

import (
	"github.com/pterm/pterm"
	"github.com/sllt/hole/internal/client/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shellCmd)
}

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "hole shell [start|stop] [agent]",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			pterm.Println("command is required")
			return
		}

		if args[0] != "start" && args[0] != "stop" {
			pterm.Println("command is not supported")
			return
		}

		var (
			result string
		)

		if args[0] == "start" {

			result, _ = service.Client.StartShell(args[1])
		} else if args[0] == "stop" {
			result, _ = service.Client.StopShell(args[1])
		}

		pterm.Println(result)
	},
}
