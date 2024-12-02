package main

import (
	"log"

	"github.com/hotaru51/dns-record-updater/config"
	"github.com/hotaru51/dns-record-updater/myip"
)

func main() {
	log.Println("load config file")
	config := config.LoadConfig()
	log.Println(config)
	log.Println(myip.GetMyIP())
}
