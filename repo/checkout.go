package repo

import (
	"MultiGit/commands"
	"MultiGit/types"
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	errorpkg "github.com/pkg/errors"
	"path/filepath"
)

func Checkout(repos *[]types.Repo, destDir string, branch string) {
	commands.ProcessReposWithProgress(repos, func(repo types.Repo) error {

		if branch == "" {
			branch = repo.Branch
		}

		repoPath := filepath.Join(destDir, repo.Name)
		failedMessage := "Checkout failed!"

		gitRepo, err := git.PlainOpen(repoPath)
		if err != nil {
			errMessage := fmt.Sprintf("Failed to open repository: %s", err)
			return errorpkg.Errorf(logFormat, repo.Name, failedMessage, errMessage)
		}

		workTree, err := gitRepo.Worktree()
		if err != nil {
			errMessage := fmt.Sprintf("Could not get worktree: %s", err)
			return errorpkg.Errorf(logFormat, repo.Name, failedMessage, errMessage)
		}

		err = workTree.Checkout(&git.CheckoutOptions{
			Branch: plumbing.NewBranchReferenceName(branch),
		})

		if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
			errMessage := fmt.Sprintf("Failed to checked out to branch: %s", err)
			return errorpkg.Errorf(logFormat, repo.Name, failedMessage, errMessage)
		}

		return nil
	}, false)
}
