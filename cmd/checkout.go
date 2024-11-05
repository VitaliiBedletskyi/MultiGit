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
	Short: "Checkout all initialized repositories according to .mgitrc config to the specified branch",
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//checkoutCmd.Flags().StringP("path", "p", "", "Path to target folder where command should be ran")
}
