// アプリケーションの設定を保持する
package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

/*
実行ファイルの絶対パスを取得する
*/
func getExecutableDirectoryPath() string {
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatalf("cannot get executable path: %s\n", err)
	}

	executableDir := filepath.Dir(executablePath)
	log.Printf("executableDir: %s\n", executableDir)

	return executableDir
}

/*
設定ファイルの存在確認を行う
*/
func isConfigFileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return err == nil
}

/*
実行ファイルと同じディレクトリかCONFIG_FILE_PATHの設定ファイルのパスを返す(どちらも存在しない場合は空文字)
*/
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

/*
YAMLファイルを読み込み内容を文字列で返す
*/
func loadYamlFile(filepath string) string {
	b, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("cannot read file: %s\n", filepath)
	}

	return string(b)
}

/*
読み込んだYAMLの文字列をパースして*Configを返す
*/
func parseYaml(text string) *Config {
	config := &Config{}

	err := yaml.Unmarshal([]byte(text), config)
	if err != nil {
		log.Fatalf("failed to parse yaml\n")
	}

	return config
}

/*
設定ファイル(YAML)を読み込み*Configを返す
*/
func LoadConfig() *Config {
	path := selectConfigFilePath()
	if len(path) <= 0 {
		log.Fatalln("config file not found")
	}

	text := loadYamlFile(path)

	return parseYaml(text)
}
