package cmd

import (
	"MultiGit/config"
	"MultiGit/log"
	"MultiGit/repo"
	"MultiGit/types"
	"fmt"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new .mgitrc configuration file in your specified directory.",
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		path, _ := cmd.Flags().GetString("path")

		mgitInitPath, err := repo.GetPath(path)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to find target folder path: %s", err))
			return
		}

		repos, err := repo.InitExistingRepos(mgitInitPath, skipRepos)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to process folder: %s: %s", mgitInitPath, err))
			return
		}

		if len(*repos) == 0 {
			log.Error("No existing repositories found")
			return
		}

		log.Table[types.Repo](*repos)

		saveErr := config.Save(mgitInitPath, types.Config{Repositories: repos}, force)
		if saveErr != nil {
			log.Error(fmt.Sprintf("Failed to save .mgit config: %s", saveErr))
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP("force", "f", false, "Force to re-initialize .mgitrc config from scratch even if the .mgitrc config is present in a managed folder")
}
