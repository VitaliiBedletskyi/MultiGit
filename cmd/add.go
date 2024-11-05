package cmd

import (
	"MultiGit/config"
	"MultiGit/log"
	"MultiGit/repo"
	"MultiGit/types"
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new repository and clone it",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Error("A repository URL must be provided.")
		}

		path, _ := cmd.Flags().GetString("path")
		repoUrl := args[0]
		repoName, err := repo.ParseRepoName(repoUrl)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to parse repository name: %s", err))
			return
		}

		mgitPath, err := repo.GetPath(path)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to find target folder path: %s", err))
			return
		}

		repository := types.Repo{
			Name:   repoName,
			URL:    repoUrl,
			Branch: "",
		}
		err = repo.Clone(&repository, mgitPath, true)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to clone repository: %s", err))
			return
		}

		err = config.Save(mgitPath, types.Config{
			Repositories: &[]types.Repo{repository},
		}, false)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to update .mgit config: %s", err))
			return
		}

		log.Success(fmt.Sprintf("Repository %s added successfully.", repository.Name))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//checkoutCmd.Flags().StringP("path", "p", "", "Path to target folder where command should be ran")
}
