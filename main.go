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
	log.Println(myip.GetMyIP())
	gc := gandi.NewClient(config.Domain, config.AccessToken)
	log.Printf("%v\n", gc)
	records, err := gc.GetRecords()
	if err == nil {
		for _, r := range records {
			log.Println(r)
		}
	}
}
