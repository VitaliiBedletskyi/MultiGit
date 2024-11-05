package repo

import (
	"MultiGit/commands"
	"MultiGit/types"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	errorpkg "github.com/pkg/errors"
	"path/filepath"
)

func Pull(repos *[]types.Repo, destDir string) {
	commands.ProcessReposWithProgress(repos, func(repo types.Repo) error {

		repoPath := filepath.Join(destDir, repo.Name)
		sshKeyPath := "/Users/bedletskyi/.ssh/id_rsa"
		failedMessage := "Pull failed!"

		publicKey, err := ssh.NewPublicKeysFromFile("git", sshKeyPath, "")
		if err != nil {
			errMessage := fmt.Sprintf("Failed to load SSH keys: %s", err)
			return errorpkg.Errorf(logFormat, repo.Name, failedMessage, errMessage)
		}

		gitRepo, err := git.PlainOpen(repoPath)
		if err != nil {
			errMessage := fmt.Sprintf("Failed to open repository: %s", err)
			return errorpkg.Errorf(logFormat, repo.Name, failedMessage, errMessage)
		}

		err = gitRepo.Fetch(&git.FetchOptions{
			RemoteName: "origin",
			Auth:       publicKey,
			Tags:       git.AllTags,
		})
		if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
			errMessage := fmt.Sprintf("Failed to fetch repository: %s", err)
			return errorpkg.Errorf(logFormat, repo.Name, failedMessage, errMessage)
		}

		workTree, err := gitRepo.Worktree()
		if err != nil {
			errMessage := fmt.Sprintf("Could not get worktree: %s", err)
			return errorpkg.Errorf(logFormat, repo.Name, failedMessage, errMessage)
		}

		err = workTree.Pull(&git.PullOptions{
			RemoteName: "origin",
			Auth:       publicKey,
		})

		if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
			errMessage := fmt.Sprintf("Failed to pull repository: %s", err)
			return errorpkg.Errorf(logFormat, repo.Name, failedMessage, errMessage)
		}

		return nil
	}, false)
}
