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
更新リクエスト用DNSレコード
*/
type DomainRecordRequest struct {
	RrsetName   string   `json:"rrset_name"`   // レコード名
	RrsetType   string   `json:"rrset_type"`   // レコードタイプ
	RrsetValues []string `json:"rrset_values"` // レコードの値
}

/*
更新リクエスト用DNSレコードのslice
*/
type DomainRecordRequestItems struct {
	Items []*DomainRecordRequest `json:"items"`
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

/*
DomainRecordResultからDomainRecordRequestを生成する
*/
func (r *DomainRecordResult) GenerateUpdateRequest(values *[]string) *DomainRecordRequest {
	return &DomainRecordRequest{
		RrsetName: r.RrsetName,
		RrsetType: r.RrsetType,
		RrsetValues: *values,
	}
}

/*
DomainRecordRequestを文字列で返す
*/
func (r *DomainRecordRequest) String() string {
	return fmt.Sprintf(
		"name: %s, type: %s, value: %v",
		r.RrsetName,
		r.RrsetType,
		r.RrsetValues,
	)
}
