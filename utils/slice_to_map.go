package utils

func SliceToMap(slice []string) map[string]bool {
	skipRepoMap := make(map[string]bool)
	for _, repoName := range slice {
		skipRepoMap[repoName] = true
	}
	return skipRepoMap
}
