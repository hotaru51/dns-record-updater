// Gnadi APIを操作するパッケージ
package gandi

import (
	"io"
	"net/http"
	"net/url"
	"encoding/json"
)

const (
	API_ENDPOINT = "https://api.gandi.net/v5/livedns/domains" // Gandi APIエンドポイント
)

/*
Gnadi API実行用のHTTPクライアント
*/
type Client struct {
	baseUrl *url.URL     // APIエンドポイントのベースURL
	domain string        // 対象ドメイン
	token string         // Gandi APIのトークン
	client *http.Client  // HTTPクライアント
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

/*
DNSレコードを取得する
	対象はAレコードのみ
*/
func (c *Client) GetRecords() ([]*DomainRecordResult, error) {
	url := c.baseUrl.JoinPath(c.domain, "records")
	q := url.Query()
	q.Add("rrset_type", "A")
	url.RawQuery = q.Encode()

	req := c.newBaseRequest("GET", url.String(), nil)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var record []*DomainRecordResult
	err = json.Unmarshal(b, &record)
	if err != nil {
		return nil, err
	}

	return record, nil
}
