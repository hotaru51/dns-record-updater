package main

import (
	"log"

	"github.com/hotaru51/dns-record-updater/config"
)

func main() {
	log.Println("load config file")
	config := config.LoadConfig()
	log.Println(config)
}
