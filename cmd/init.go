package cmd

import (
	"github-addon-backup/installations"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize github repo for addon and accunts with no git repo",
	Run: func(cmd *cobra.Command, args []string) {
		installations.Init(baseWowDir)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
