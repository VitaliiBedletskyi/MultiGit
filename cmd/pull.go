package cmd

import (
	"MultiGit/config"
	"MultiGit/log"
	"MultiGit/repo"
	"fmt"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull all initialized repositories according to .mgitrc config",
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

		repo.Pull(mgitConfig.Repositories, mgitPath)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//pullCmd.Flags().BoolP("force", "f", false, "Force the repository to be cloned even if a target folder isn't empty. All data in a target folder will be lost.")
}
