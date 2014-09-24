package main

import (
	"flag"
	"log"
	"time"

	"github.com/longnguyen11288/lazy-collector/cleaner"
	"github.com/longnguyen11288/lazy-collector/config"
	"github.com/longnguyen11288/lazy-collector/downloader"
	"github.com/longnguyen11288/lazy-collector/extractor"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "c", "", "config file")
	flag.Parse()
	collectorConfig, err := config.ParseConfigFile(configFile)
	if err != nil {
		log.Fatalf("Unable to parse config: %s\n", err)
	}
	go downloader.Run(collectorConfig)
	go extractor.Run(collectorConfig)
	go cleaner.Run(collectorConfig)
	for {
		time.Sleep(time.Duration(collectorConfig.Sleep) * time.Minute)
	}
}
