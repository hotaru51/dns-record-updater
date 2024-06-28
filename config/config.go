package config

type Config struct {
	Domain string		`yaml:"domain"`
	AccessToken string	`yaml:"access_token"`
	Records []string	`yaml:"records"`
}
}

const (
	API_ENDPOINT = "https://api.gandi.net/v5/livedns/domains"
	CONFIG_FILE_PATH = "/etc/htrap/dns-record-updater"
	CONFIG_FILE_NAME = "config.yaml"
)
