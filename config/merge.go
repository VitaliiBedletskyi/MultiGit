package config

import "MultiGit/types"

func mergeRepos(repos1, repos2 *[]types.Repo) *[]types.Repo {
	// Create a map to store unique repos by name
	repoMap := make(map[string]types.Repo)

	// Add all repos from the first slice to the map
	for _, repo := range *repos1 {
		repoMap[repo.Name] = repo
	}

	// Add repos from the second slice, replacing if the name already exists
	for _, repo := range *repos2 {
		repoMap[repo.Name] = repo
	}

	// Collect merged repos into a slice
	mergedRepos := make([]types.Repo, 0, len(repoMap))
	for _, repo := range repoMap {
		mergedRepos = append(mergedRepos, repo)
	}

	return &mergedRepos
}
