// Gnadi APIを操作するパッケージ
package gandi

import (
	"io"
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
	domain string
	token string
	client *http.Client
}

/*
クライアントを生成する
*/
func NewClient(domain string, token string) *Client {
	u, _ := url.Parse(API_ENDPOINT)
	c := &Client{
		baseUrl: u,
		domain: domain,
		token: token,
		client: &http.Client{},
	}

	return c
}

/*
トークン設定済みのリクエストのベースを生成
*/
func (c *Client) newBaseRequest(method string, url string, body io.Reader) *http.Request {
	r, _ := http.NewRequest(method, url, body)
	r.Header.Add("Authorization", "Bearer " + c.token)

	return r
}
