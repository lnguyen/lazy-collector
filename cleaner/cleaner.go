package cleaner

import (
	"time"

	"github.com/alexcesaro/log"
	"github.com/longnguyen11288/lazy-collector/config"
)

//Log is logger
var Log log.Logger

//Run some stuff yo
func Run(collectorConfig *config.Config) {
	Log = collectorConfig.Log
	Log.Info("Starting Cleaner")
	for {
		Log.Debug("Starting Cleaner run")
		torrents, err := collectorConfig.TransmissionClient.GetTorrents()
		if err != nil {
			Log.Errorf("Error getting torrents %v", err)
		}
		for _, torrent := range torrents {
			if torrent.UploadRatio > collectorConfig.SeedRatio {
				collectorConfig.TransmissionClient.RemoveTorrent(torrent.ID, true)
				Log.Infof("Removing %s", torrent.Name)
			}
		}
		Log.Debug("Finished Cleaner run")
		time.Sleep(3 * time.Minute)
	}
}
