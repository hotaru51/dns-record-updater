package config

import (
	"log"
	"os"
	"path/filepath"
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
