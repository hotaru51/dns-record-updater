package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func getExecutableDirectoryPath() string {
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatalf("cannot get executable path: %s\n", err)
	}

	executableDir := filepath.Dir(executablePath)
	log.Printf("executableDir: %s\n", executableDir)

	return executableDir
}

func isConfigFileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return err == nil
}

func selectConfigFilePath() string {
	localConfigPath := getExecutableDirectoryPath() + "/" + CONFIG_FILE_NAME
	if isConfigFileExists(localConfigPath) {
		return localConfigPath
	}

	globalConfigPath := CONFIG_FILE_PATH + "/" + CONFIG_FILE_PATH
	if isConfigFileExists(globalConfigPath) {
		return globalConfigPath
	}

	return ""
}

func loadYamlFile(filepath string) string {
	b, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("cannot read file: %s\n", filepath)
	}

	return string(b)
}

func parseYaml(text string) *Config {
	config := &Config{}

	err := yaml.Unmarshal([]byte(text), config)
	if err != nil {
		log.Fatalf("failed to parse yaml\n")
	}

	return config
}
