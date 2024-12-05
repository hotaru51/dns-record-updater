// 自身のグローバルIPアドレスを確認するパッケージ
package myip

import (
	"net/http"
	"io"
	"strings"
)

const (
	CHECK_IP_URL = "https://checkip.amazonaws.com/" // IPアドレス確認先URL
)

/*
自身のグローバルIPを返す
*/
func GetMyIP() string {
	res, err := http.Get(CHECK_IP_URL)
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return ""
	}

	return strings.Trim(string(b), "\n")
}
