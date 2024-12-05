// アプリケーションの設定を保持する
package config

import (
	"fmt"
)

/*
プリケーションの設定値を保持
*/
type Config struct {
	Domain string		`yaml:"domain"` // 設定対象のドメイン
	AccessToken string	`yaml:"access_token"` // Gandi APIのトークン
	Records []string	`yaml:"records"` // 更新対象のDNSレコード
}

/*
Configを文字列で返す
*/
func (c *Config) String() string {
	return fmt.Sprintf("domain: %s, access_token: ***, records: %v", c.Domain, c.Records)
}

const (
	CONFIG_FILE_PATH = "/etc/htrap/dns-record-updater" // 設定ファイル配置場所
	CONFIG_FILE_NAME = "config.yaml" // 設定ファイル名
)
