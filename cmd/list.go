package cmd

import (
	"log"

	"github.com/greendwin/l2chat/methods"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List network devices",
	Long:  "List available network devices with their ID and associated IP addresses.",
	Run: func(cmd *cobra.Command, args []string) {
		err := methods.ListDevices()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
