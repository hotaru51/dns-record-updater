package config

type Config struct {
	Domain string
	AccessToken string
	Records []string
}

const (
	API_ENDPOINT = "https://api.gandi.net/v5/livedns/domains"
	CONFIG_FILE_PATH = "/etc/htrap/dns-record-updater"
	CONFIG_FILE_NAME = "config.yaml"
)
