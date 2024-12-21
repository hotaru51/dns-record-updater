package main

import (
	"log"
	"slices"

	"github.com/hotaru51/dns-record-updater/config"
	"github.com/hotaru51/dns-record-updater/myip"
	"github.com/hotaru51/dns-record-updater/gandi"
)

func main() {
	log.Println("load config file")
	config := config.LoadConfig()
	log.Printf("config file loaded: %s\n", config)

	log.Println("get current global IP")
	currentIp := myip.GetMyIP()
	log.Printf("current global ip: %s\n", currentIp)

	gc := gandi.NewClient(config.Domain, config.AccessToken)

	log.Println("fetch currnet A records")
	records, err := gc.GetRecords()
	if err != nil {
		log.Fatalf("fetch A records failed: %s\n", err)
	}
	log.Printf("num of records: %d\n", len(records))

	var target []string

	// 設定に新規レコードが含まれるか確認
	for _, c := range config.Records {
		isNew := !slices.ContainsFunc(records, func(r *gandi.DomainRecordResult) bool {
			return r.RrsetName == c
		})

		if isNew {
			log.Printf("%s is a new record, so it will be created", c)
			target = append(target, c)
		}
	}

	// configのRecordsに含まれ、更新対象であるレコードを確認
	for _, r := range records {
		isContain := slices.ContainsFunc(config.Records, func(c string) bool {
			return c == r.RrsetName
		})
		
		if isContain && r.ShouldUpdate(currentIp) {
			log.Printf("%s will be updated (prev: %v)", r.RrsetName, r.RrsetValues)
			target = append(target, r.RrsetName)
		}
	}

	if len(target) <= 0 {
		log.Println("no records to update")
		return
	}

	log.Println("update target records")
	for _, t := range target {
		log.Printf("%s: update\n", t)
		req := gandi.NewDomainRecordRequestItems(currentIp)
		res, err := gc.UpdateRecord(t, req)
		if err != nil {
			log.Printf("%s: update failed: %s\n", t, err)
		} else {
			log.Printf("%s: result: %s\n", t, res)
		}
	}
}
