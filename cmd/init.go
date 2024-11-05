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
	Short: "Init the .mgitrc config file in the current folder",
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		path, _ := cmd.Flags().GetString("path")

		mgitInitPath, err := repo.GetPath(path)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to find target folder path: %s", err))
			return
		}

		repos, err := repo.InitExistingRepos(mgitInitPath)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to read .mgit config: %s", err))
			return
		}

		if len(*repos) == 0 {
			log.Error(fmt.Sprintf("Failed to read .mgit config: %s", err))
			fmt.Println("No existing repos found")
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().BoolP("force", "f", false, "Force to re-initialize .mgitrc config from scratch even if the .mgitrc config is present in a managed folder")
}
