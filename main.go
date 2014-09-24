package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/longnguyen11288/lazy-collector/config"
	"github.com/longnguyen11288/lazy-collector/downloader"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "c", "", "config file")
	flag.Parse()
	collectorConfig, err := config.ParseConfigFile(configFile)
	if err != nil {
		log.Fatalf("Unable to parse config: %s\n", err)
	}
	fmt.Println(collectorConfig)

	collectorConfig.Log.Info("Parsed config")
	go downloader.Run(collectorConfig)
	collectorConfig.Log.Info("Ran downloader")
	for {
		time.Sleep(5 * time.Minute)
	}
}
