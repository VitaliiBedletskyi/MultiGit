package repo

import (
	"MultiGit/commands"
	"MultiGit/types"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/pkg/errors"
	"path/filepath"
)

const logFormat = "%-25s %-20s %s"

func Clone(repos *[]types.Repo, destDir string) {

	commands.ProcessReposWithProgress(repos, func(repo types.Repo) error {
		// Destination folder for each repo
		repoPath := filepath.Join(destDir, repo.Name)
		branch := plumbing.NewBranchReferenceName(repo.Branch)
		sshKeyPath := "/Users/bedletskyi/.ssh/id_rsa"

		publicKey, err := ssh.NewPublicKeysFromFile("git", sshKeyPath, "")
		if err != nil {
			errMessage := fmt.Sprintf("Could not load SSH keys: %s", err)
			return errors.Errorf(logFormat, repo.Name, "Clone failed!", errMessage)
		}

		clonedRepo, cloneErr := git.PlainClone(repoPath, false, &git.CloneOptions{
			URL:           repo.URL,
			Auth:          publicKey,
			ReferenceName: branch,
		})

		if cloneErr != nil {
			errMessage := fmt.Sprintf("Could not clone repo: %s", cloneErr)
			return errors.Errorf(logFormat, repo.Name, "Clone failed!", errMessage)
		}

		originHeadRef := plumbing.NewSymbolicReference(
			"refs/remotes/origin/HEAD",
			plumbing.NewRemoteReferenceName("origin", repo.Branch),
		)

		err = clonedRepo.Storer.SetReference(originHeadRef)
		if err != nil {
			errMessage := fmt.Sprintf("Could not create symbolic reference: %s", err)
			return errors.Errorf(logFormat, repo.Name, "Clone failed!", errMessage)
		}

		return nil
	}, false)
}
