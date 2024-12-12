package main

import (
	"log"

	"github.com/hotaru51/dns-record-updater/config"
	"github.com/hotaru51/dns-record-updater/myip"
	"github.com/hotaru51/dns-record-updater/gandi"
)

func main() {
	log.Println("load config file")
	config := config.LoadConfig()
	log.Println(config)
	ip := myip.GetMyIP()
	log.Println(ip)
	gc := gandi.NewClient(config.Domain, config.AccessToken)
	log.Printf("%v\n", gc)
	records, err := gc.GetRecords()
	if err == nil {
		log.Println("current records")
		for _, r := range records {
			log.Println(r)
		}
	}
	log.Println("generate update record request")
	for _, r := range records {
		dreq := r.GenerateUpdateRequest(&[]string{ip})
		log.Println(dreq)
	}
}
