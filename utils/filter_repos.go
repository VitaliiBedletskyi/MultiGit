package utils

import "MultiGit/types"

func FilterRepos(repos *[]types.Repo, skipRepos []string) *[]types.Repo {
	skipRepoMap := SliceToMap(skipRepos)
	var result []types.Repo
	for _, repo := range *repos {
		if !skipRepoMap[repo.Name] {
			result = append(result, repo)
		}
	}
	return &result
}
