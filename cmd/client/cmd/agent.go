package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"hole/internal/client/service"
)

func init() {
	rootCmd.AddCommand(agentCmd)
}

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "hole agent",
	Run: func(cmd *cobra.Command, args []string) {

		list, err := service.Client.GetAgentList()
		if err != nil {
			pterm.Println(err)
			return
		}

		tableData1 := pterm.TableData{
			{"ShorName", "HostName", "OS", "Description"},
		}

		for _, v := range list {
			tableData1 = append(tableData1,
				[]string{v.ShortName, v.Hostname, v.OS, v.Description})
		}

		// Create a table with a header and the defined data, then render it
		pterm.DefaultTable.WithHasHeader().WithData(tableData1).Render()

		pterm.Println()
	},
}
