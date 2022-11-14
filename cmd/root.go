/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github-addon-backup/config"
	"github-addon-backup/installations"
	"github.com/spf13/cobra"
	"log"
)

var pull bool
var baseWowDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-addon-backup",
	Short: "Manage wow addons in a public github repository",
	Run: func(cmd *cobra.Command, args []string) {
		if pull {
			installations.KeepUpToDate(baseWowDir)
		} else {

		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	defaultPath := config.DefaultPathAndMessage()

	rootCmd.Flags().StringVarP(&baseWowDir, "base-wow-directory", "d", defaultPath, "Base directory for World of Warcraft installation.")
	rootCmd.Flags().BoolVarP(&pull, "pull", "p", false, "Help message for toggle")
}


