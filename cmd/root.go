/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github-addon-backup/config"
	"github.com/spf13/cobra"
	"log"
)

var pull bool
var baseWowDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-addon-backup",
	Short: "Manage wow addons in a public github repository",
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
}
