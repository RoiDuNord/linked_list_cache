package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getConfigFilePath() (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %v", err)
	}

	rootDir, err := trimToRoot(workingDir)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	configFilePath := filepath.Join(rootDir, "config", "config.yaml")

	configFilePath = filepath.Clean(configFilePath)

	return configFilePath, nil
}

func trimToRoot(workingDir string) (string, error) {
	targetDir := "workWithCache"
	index := strings.Index(workingDir, targetDir)
	if index == -1 {
		return "", fmt.Errorf("filepath (%s) does not contain 'workWithCache' directory", workingDir)
	}
	fmt.Printf("index: %d\n", index)
	rootDirs := workingDir[:index]
	pathToConfig := rootDirs + targetDir
	return pathToConfig, nil
}
