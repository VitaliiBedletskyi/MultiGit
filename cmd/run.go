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
	Short: "Run specified command in all initialized repositories according to .mgitrc config",
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//runCmd.Flags().BoolP("force", "f", false, "Force the repository to be cloned even if a target folder isn't empty. All data in a target folder will be lost.")
}
