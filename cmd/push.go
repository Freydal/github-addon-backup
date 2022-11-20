package cmd

import (
	"github-addon-backup/installations"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push configured repos",
	Run: func(cmd *cobra.Command, args []string) {
		installations.UpdateDaily(baseWowDir)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
