package downloader

import (
	"time"

	"github.com/alexcesaro/log"
	"github.com/longnguyen11288/lazy-collector/config"
)

var shows Shows

//Log is logger
var Log log.Logger

//Run starts downloader
func Run(collectorConfig *config.Config) {
	Log = collectorConfig.Log
	Log.Info("Starting downloader")

	shows = collectorConfig.Shows
	Log.Debugf("Shows %s", shows)
	for {
		Log.Debug("Started Downloader run")
		Log.Debug("Getting feeds")
		feedItems, err := Feeds(collectorConfig.Rss)
		if err != nil {
			Log.Errorf("Error collecting feeds %s", err)
		}
		for _, item := range feedItems {
			collectorConfig.Log.Debugf("Feed item %s", item.title)
			isDownloadable, show := shows.IsDownloadableShow(item.title)
			if isDownloadable {
				item.Show = show
				item.TimeAdded = time.Now()
				item.GetShowData()
				if !Blacklisted(item) {
					Log.Infof("Adding torrent %s", item.title)
					_, err = collectorConfig.TransmissionClient.AddTorrentByURL(item.Link, collectorConfig.DownloadDir)
					if err != nil {
						Log.Errorf("Error adding torrent %s", err)
					}
					AddToBlackList(item)
				}
			}
		}
		Log.Debug("Finished Downloader run")
		time.Sleep(time.Duration(collectorConfig.Sleep) * time.Minute)
	}

}
