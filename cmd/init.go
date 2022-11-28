package cmd

import (
	"github-addon-backup/gh"
	"github-addon-backup/installations"
	"github.com/cli/cli/v2/pkg/cmdutil"
	"github.com/spf13/cobra"
	"log"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize github repo for addon and accunts with no git repo",
	Run: func(cmd *cobra.Command, args []string) {
		err := gh.GetStatus()
		if err == cmdutil.SilentError {
			err := gh.Login()
			if err != nil {
				log.Fatal("failed to login gh", err)
			}
		} else if err != nil {
			log.Fatal("Failed to get stats of gh", err)
		}
		installations.Init(baseWowDir)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
