package config

import (
	"MultiGit/types"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

func Read(path string) (*types.Config, error) {
	if !strings.HasSuffix(path, ".mgitrc") {
		path = filepath.Join(path, ".mgitrc")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config types.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
