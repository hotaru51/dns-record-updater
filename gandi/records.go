// Gnadi APIを操作するパッケージ
package gandi

/*
DNSレコード
*/
type Record struct {
	RrsetName   string   `json:"rrset_name"`   // レコード名
	RrsetTTL    int      `json:"rrset_ttl"`    // TTL
	RrsetType   string   `json:"rrset_type"`   // レコードタイプ
	RrsetValues []string `json:"rrset_values"` // レコードの値
	RrsetHref   string   `json:"rrset_href"`   // レコードのURL
}
