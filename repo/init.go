package repo

import (
	"MultiGit/commands"
	"MultiGit/types"
	"MultiGit/utils"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var systemFolders = map[string]bool{
	".idea":   true,
	".vscode": true,
	".git":    true,
}

func InitExistingRepos(folderPath string, skipRepos []string) (*[]types.Repo, error) {
	var wg sync.WaitGroup

	repos := make([]types.Repo, 0)
	errChan := make(chan error, 100)
	doneChan := make(chan types.Repo, 100)

	files, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	skipReposMap := utils.SliceToMap(skipRepos)

	for _, file := range files {
		if (file.IsDir() && !systemFolders[file.Name()]) || !skipReposMap[file.Name()] {
			getGitRepoInfo(file.Name(), folderPath, &wg, doneChan, errChan)
		}
	}

	go func() {
		wg.Wait()
		close(doneChan)
		close(errChan)
	}()

	for doneChan != nil || errChan != nil {
		select {
		case repo, ok := <-doneChan:
			if !ok {
				// `doneChan` is closed - exit the loop
				doneChan = nil
				continue
			}

			repos = append(repos, repo)
		case err, ok := <-errChan:
			if !ok {
				// `errChan` is closed - exit the loop
				errChan = nil
				continue
			}

			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return &repos, nil
}

func getRepoInfoError(repoName string, err error) error {
	return fmt.Errorf("the %s folder: %s", repoName, err.Error())
}

func getGitRepoInfo(repoName string, path string, wg *sync.WaitGroup, doneChan chan<- types.Repo, errChan chan<- error) {
	wg.Add(1)
	go func(repoPath string) {
		defer wg.Done()

		// Open the repository (specify the path to your repo)
		repo, err := git.PlainOpen(repoPath)
		if err != nil {
			errChan <- getRepoInfoError(repoName, err)
			return
		}

		remote, err := repo.Remote("origin")
		if err != nil {
			errChan <- getRepoInfoError(repoName, err)
			return
		}

		remoteURL, err := getRemoteURL(remote)
		if err != nil {
			errChan <- getRepoInfoError(repoName, err)
		}

		defaultBranch, err := getDefaultBranch(repoPath)
		if err != nil {
			errChan <- getRepoInfoError(repoName, err)
			return
		}

		// Send the Repo information to doneChan
		doneChan <- types.Repo{Name: repoName, URL: remoteURL, Branch: defaultBranch}
	}(filepath.Join(path, repoName))
}

func getRemoteURL(remote *git.Remote) (string, error) {
	remoteConfig := remote.Config()
	if len(remoteConfig.URLs) == 0 {
		return "", errors.New("no URLs found for remote")
	}
	return remoteConfig.URLs[0], nil
}

func getDefaultBranch(repoPath string) (string, error) {
	fullBranchName, err := commands.RunGitCommand(repoPath, "symbolic-ref", "--short", "refs/remotes/origin/HEAD")
	if err != nil {
		return "", err
	}

	// Process the output to get only the branch name
	defaultBranch := strings.TrimPrefix(fullBranchName, "origin/")
	return strings.TrimSpace(defaultBranch), nil
}
