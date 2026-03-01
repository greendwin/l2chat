package cmd

import (
	"log"

	"github.com/greendwin/l2chat/methods"
	"github.com/spf13/cobra"
)

var devices []string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run L2Chat server",
	Long:  `Run L2Chat server on specified devices with provided name.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(devices) == 0 {
			log.Fatal("at least one `--device` must be provided")
		}

		err := methods.RunServer(args[0], devices)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.ArgAliases = append(runCmd.ArgAliases, "name")
	runCmd.Flags().StringArrayVarP(&devices, "device", "e", nil, "Network devices to run on")
}
