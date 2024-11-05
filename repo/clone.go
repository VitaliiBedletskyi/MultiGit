package repo

import (
	"MultiGit/commands"
	"MultiGit/types"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

const logFormat = "%-25s %-20s %s"

// ParseRepoName extracts the repository name from a given URL (HTTP or SSH).
func ParseRepoName(url string) (string, error) {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		// Example: http://github.com/user/repo.git or https://github.com/user/repo.git
		parts := strings.Split(url, "/")
		repoWithExtension := parts[len(parts)-1]
		return strings.TrimSuffix(repoWithExtension, ".git"), nil
	} else if strings.HasPrefix(url, "git@") {
		// Example: git@github.com:user/repo.git
		parts := strings.Split(url, ":")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid SSH URL format")
		}
		repoParts := strings.Split(parts[1], "/")
		repoWithExtension := repoParts[len(repoParts)-1]
		return strings.TrimSuffix(repoWithExtension, ".git"), nil
	} else {
		return "", fmt.Errorf("unsupported URL format")
	}
}

func Clone(repo *types.Repo, destDir string, showProgress bool) error {
	// Destination folder for each repo
	repoPath := filepath.Join(destDir, repo.Name)
	failedCloneMesssage := "Clone failed!"
	sshKeyPath := "/Users/bedletskyi/.ssh/id_rsa"

	publicKey, err := ssh.NewPublicKeysFromFile("git", sshKeyPath, "")
	if err != nil {
		errMessage := fmt.Sprintf("Failed to load SSH keys: %s", err)
		return errors.Errorf(logFormat, repo.Name, failedCloneMesssage, errMessage)
	}

	cloneOptions := &git.CloneOptions{
		URL:  repo.URL,
		Auth: publicKey,
	}

	if showProgress {
		cloneOptions.Progress = os.Stdout
	}

	if repo.Branch != "" {
		cloneOptions.ReferenceName = plumbing.NewBranchReferenceName(repo.Branch)
	}

	clonedRepo, cloneErr := git.PlainClone(repoPath, false, cloneOptions)

	if cloneErr != nil {
		errMessage := fmt.Sprintf("Failed to clone repo: %s", cloneErr)
		return errors.Errorf(logFormat, repo.Name, failedCloneMesssage, errMessage)
	}

	if repo.Branch == "" {
		// Retrieving the remote default branch
		ref, err := clonedRepo.Reference(plumbing.HEAD, true)
		if err != nil {
			errMessage := fmt.Sprintf("Failed to get repository default branch: %s", err)
			return errors.Errorf(logFormat, repo.Name, failedCloneMesssage, errMessage)
		}

		repo.Branch = ref.Name().Short()
	}

	originHeadRef := plumbing.NewSymbolicReference(
		"refs/remotes/origin/HEAD",
		plumbing.NewRemoteReferenceName("origin", repo.Branch),
	)

	err = clonedRepo.Storer.SetReference(originHeadRef)
	if err != nil {
		errMessage := fmt.Sprintf("Failed to create symbolic reference: %s", err)
		return errors.Errorf(logFormat, repo.Name, failedCloneMesssage, errMessage)
	}

	return nil
}

func CloneAll(repos *[]types.Repo, destDir string) {
	commands.ProcessReposWithProgress(repos, func(repo types.Repo) error {
		return Clone(&repo, destDir, false)
	}, false)
}
