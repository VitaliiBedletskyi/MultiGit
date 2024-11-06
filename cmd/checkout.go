package cmd

import (
	"MultiGit/config"
	"MultiGit/log"
	"MultiGit/repo"
	"MultiGit/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Switch all repositories to the specified branch. If no branch is specified, will branch from .mgitrc config.",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		branch := ""
		if len(args) > 0 {
			branch = args[0]
		}

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

		repo.Checkout(utils.FilterRepos(mgitConfig.Repositories, skipRepos), mgitPath, branch)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
