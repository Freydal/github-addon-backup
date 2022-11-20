package cmd

import (
	"github-addon-backup/installations"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull configured git repos",
	Run: func(cmd *cobra.Command, args []string) {
		installations.KeepUpToDate(baseWowDir)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
