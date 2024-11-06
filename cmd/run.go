package cmd

import (
	"MultiGit/config"
	"MultiGit/log"
	"MultiGit/repo"
	"MultiGit/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute any custom command across all repositories, supporting batch execution of scripts or Git commands.",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")

		if len(args) == 0 || args[0] == "" {
			log.Error("No command specified")
			return
		}

		command := args[0]

		mgitPath, err := repo.GetPath(path)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to find target folder path: %s", err))
			return
		}

		mgitConfig, err := config.Read(mgitPath)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to read .mgit config: %s", err))
			return
		}

		repo.Run(utils.FilterRepos(mgitConfig.Repositories, skipRepos), mgitPath, command)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
