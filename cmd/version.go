package cmd

import (
	"clibootstrap/globals"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of clibootstrap",
	Run: func(cmd *cobra.Command, args []string) {
		globals.SugaredLogger.Infof("clibootstrap v%s", "0.01")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
