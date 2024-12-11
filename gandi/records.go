// Gnadi APIを操作するパッケージ
package gandi

import (
	"fmt"
)

/*
DNSレコード
*/
type DomainRecordResult struct {
	RrsetName   string   `json:"rrset_name"`   // レコード名
	RrsetTTL    int      `json:"rrset_ttl"`    // TTL
	RrsetType   string   `json:"rrset_type"`   // レコードタイプ
	RrsetValues []string `json:"rrset_values"` // レコードの値
	RrsetHref   string   `json:"rrset_href"`   // レコードのURL
}

/*
DomainRecordResultを文字列で返す
*/
func (r *DomainRecordResult) String() string {
	return fmt.Sprintf(
		"name: %s, TTL: %d, type: %s, value: %v",
		r.RrsetName,
		r.RrsetTTL,
		r.RrsetType,
		r.RrsetValues,
	)
}
