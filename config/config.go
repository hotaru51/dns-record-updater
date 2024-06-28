package config

import (
	"fmt"
)

type Config struct {
	Domain string		`yaml:"domain"`
	AccessToken string	`yaml:"access_token"`
	Records []string	`yaml:"records"`
}

func (c *Config) String() string {
	return fmt.Sprintf("domain: %s, access_token: ***, records: %v", c.Domain, c.Records)
}

const (
	API_ENDPOINT = "https://api.gandi.net/v5/livedns/domains"
	CONFIG_FILE_PATH = "/etc/htrap/dns-record-updater"
	CONFIG_FILE_NAME = "config.yaml"
)
