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
	RrsetType   string   `json:"rrset_type"`   // レコードタイプ
	RrsetValues []string `json:"rrset_values"` // レコードの値
	RrsetTTL    int      `json:"rrset_ttl"`    // TTL
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
レコードの更新が必要か確認する
*/
func (r *DomainRecordResult) ShouldUpdate(currentIp string) bool {
	// レコードが2以上の場合は更新
	if len(r.RrsetValues) >= 2 {
		return true
	}

	// レコードがない場合は更新するので初期値をtrueとする
	res := true
	for _, v := range r.RrsetValues {
		res = (v != currentIp)
	}

	return res
}

/*
DomainRecordRequestItemsを生成する
*/
func NewDomainRecordRequestItems(ipv4Value string) *DomainRecordRequestItems {
	return &DomainRecordRequestItems{
		Items: []*DomainRecordRequest{
			{
				RrsetType: "A",
				RrsetValues: []string{ipv4Value},
				RrsetTTL: 300,
			},
		},
	}
}

/*
DomainRecordRequestを文字列で返す
*/
func (r *DomainRecordRequest) String() string {
	return fmt.Sprintf(
		"type: %s, ttl: %d, value: %v",
		r.RrsetType,
		r.RrsetTTL,
		r.RrsetValues,
	)
}
