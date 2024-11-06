package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	skipRepos []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mgit",
	Short: "MultiGit - Effortless Git Repositories Management Tool",
	Long: `MultiGit is a simple CLI tool designed to streamline and automate batch processing of commands 
across multiple Git repositories. Ideal for developers and teams managing multiple repositories, 
MultiGit provides an intuitive way to handle various Git actions, saving you time and reducing repetitive tasks.`,
	Version: "v0.0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("path", "p", "", "Path to target folder where command should be ran")
	rootCmd.PersistentFlags().StringSliceVarP(&skipRepos, "skip", "s", []string{}, "List of repositories that will be skipped for processing")
}
