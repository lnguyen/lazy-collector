package extractor

import (
	"os/exec"
	"path/filepath"
	"time"

	"github.com/alexcesaro/log"
	"github.com/longnguyen11288/lazy-collector/config"
)

//Log is logger
var Log log.Logger

//Run some stuff yo
func Run(collectorConfig *config.Config) {
	Log = collectorConfig.Log
	Log.Info("Starting Extractor")
	for {
		Log.Debug("Starting Extractor run")
		torrents, err := collectorConfig.TransmissionClient.GetTorrents()
		if err != nil {
			Log.Errorf("Error getting torrents %v", err)
		}
		for _, torrent := range torrents {
			if torrent.PercentDone == 1 {
				torrentPath := filepath.Join(torrent.DownloadDir, torrent.Name) + "/*.rar"
				_, err := exec.Command("unrar", "e", "-y", torrentPath, collectorConfig.OutputDir).Output()
				if err != nil {
					Log.Errorf("Error running extract command %s", err)
				}
				Log.Infof("Extracting %s", torrent.Name)
			}
		}
		Log.Debug("Finished Extractor run")
		time.Sleep(time.Duration(collectorConfig.Sleep) * time.Minute)
	}
}
