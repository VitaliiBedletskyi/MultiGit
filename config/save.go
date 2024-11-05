package config

import (
	"MultiGit/types"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func Save(multiRepoPath string, config types.Config, forceSet bool) error {
	appConfigPath := filepath.Join(multiRepoPath, ".mgitrc")
	if forceSet {
		return writeConfig(appConfigPath, config)
	}

	_, existErr := os.Stat(appConfigPath)
	oldConfig, readErr := Read(appConfigPath)
	if existErr != nil || readErr != nil {
		return writeConfig(appConfigPath, config)
	}

	config.Repositories = mergeRepos(oldConfig.Repositories, config.Repositories)
	return writeConfig(appConfigPath, config)
}

func writeConfig(appConfigPath string, config types.Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	// Write the YAML data to the specified file
	err = os.WriteFile(appConfigPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
