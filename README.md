# dns-record-updater

GandiのLiveDNSに登録してるDNSレコードを現在のグローバルIPに更新するやつ  
現在はAレコードのみ対応

* 現在のグローバルIPを取得
* 設定ファイルに記載されているレコードの値から変わっていたら更新する
* レコードが存在しない場合は新規作成する

## ビルド

```sh
make
```

## 設定

* Gandiの[Personal Access Token(PAT)](https://api.gandi.net/docs/authentication/)を下記の権限で発行しておく
    * ドメイン名を確認、更新
    * ドメイン名の設定管理
* `config.sample.yaml` を `config.yaml` でコピーし、必要な設定を行う
    ```yaml
    domain: example.com # ドメイン名
    access_token: PERSONAL_ACCESS_TOKEN # 取得したトークン
    records:
      - rrset_name # レコード名
      - rrset_name
    ```
* 設定ファイルを下記のいずれかに配置する
    * 実行ファイルと同じディレクトリ
    * `/etc/htrap/dns-record-updater`

## 実行

引数、オプション等は不要

```sh
dns-record-updater
```
