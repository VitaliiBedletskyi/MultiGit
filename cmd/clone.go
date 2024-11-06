package cmd

import (
	"MultiGit/config"
	"MultiGit/log"
	"MultiGit/repo"
	"MultiGit/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone all repositories listed in your .mgitrc file.",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
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

		repo.CloneAll(utils.FilterRepos(mgitConfig.Repositories, skipRepos), mgitPath)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().BoolP("force", "f", false, "Force the repository to be cloned even if a target folder isn't empty. All data in a target folder will be lost.")
}
