package config

type Config struct {
	Domain string
	AccessToken string
	Records []string
}

const (
	API_ENDPOINT = "https://api.gandi.net/v5/livedns/domains"
)
