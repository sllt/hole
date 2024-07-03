package cmd

import (
	"github.com/pterm/pterm"
	"github.com/sllt/hole/internal/client/service"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rootCmd.AddCommand(socksCmd)
}

var socksCmd = &cobra.Command{
	Use:   "socks",
	Short: "hole exec [start|stop] [agent] [port]",
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

			port, err := strconv.Atoi(args[2])
			if err != nil {
				pterm.Println(err)
				return
			}

			result, err = service.Client.StartSocks(args[1], port)
		} else if args[0] == "stop" {
			result, _ = service.Client.StopSocks(args[1])
		}

		pterm.Println(result)
	},
}
