package repo

import (
	"MultiGit/commands"
	"MultiGit/types"
	"fmt"
	errorpkg "github.com/pkg/errors"
	"path/filepath"
)

func isOptimizedNpmInstallCommand(command string) bool {
	optimizedCommands := []string{
		"npm install", "npm i", "npm in", "npm ins", "npm inst", "npm insta", "npm instal", "npm isnt", "npm isnta",
		"npm isntal", "npm isntall", "npm ci", "npm clean-install", "npm ic", "npm install-clean", "npm isntall-clean",
	}
	for _, cmd := range optimizedCommands {
		if command == cmd {
			return true
		}
	}
	return false
}

func isAllNpmModulesInstalled(repoPath string) bool {
	_, err := commands.RunCustomCommand(repoPath, "npm ls --depth=0")
	return err == nil
}

func Run(repos *[]types.Repo, destDir string, command string) {
	fmt.Println("Running command...")

	isOptimizedCommand := isOptimizedNpmInstallCommand(command)

	commands.ProcessReposWithProgress(repos, func(repo types.Repo) error {

		repoPath := filepath.Join(destDir, repo.Name)
		failedMessage := "Command failed!"

		if isOptimizedCommand && isAllNpmModulesInstalled(repoPath) {
			return nil
		}

		_, err := commands.RunCustomCommand(repoPath, command)
		if err != nil {
			errMessage := fmt.Sprintf("Failed to run command: %s", err)
			return errorpkg.Errorf(logFormat, repo.Name, failedMessage, errMessage)
		}

		return nil
	}, true)
}
