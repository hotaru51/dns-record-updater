// Gnadi APIを操作するパッケージ
package gandi

import (
	"net/http"
	"net/url"
)

const (
	API_ENDPOINT = "https://api.gandi.net/v5/livedns/domains" // Gandi APIエンドポイント
)

/*
Gnadi API実行用のHTTPクライアント
*/
type Client struct {
	baseUrl *url.URL
	token string
	client *http.Client
}

/*
クライアントを生成する
*/
func NewClient(token string) *Client {
	u, _ := url.Parse(API_ENDPOINT)
	c := &Client{
		baseUrl: u,
		token: token,
		client: &http.Client{},
	}

	return c
}
