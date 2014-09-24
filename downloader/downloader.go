package downloader

import (
	"time"

	"github.com/longnguyen11288/lazy-collector/config"
)

var shows Shows

//Run starts downloader
func Run(collectorConfig *config.Config) {
	collectorConfig.Log.Info("Starting downloader")

	shows = collectorConfig.Shows
	collectorConfig.Log.Debugf("Shows %s\n", shows)
	for {
		collectorConfig.Log.Debug("Getting feeds")
		feedItems, err := Feeds(collectorConfig.Rss)
		if err != nil {
			collectorConfig.Log.Errorf("Error collecting feeds %s", err)
		}
		for _, item := range feedItems {
			collectorConfig.Log.Debugf("Feed item %s", item.title)
			isDownloadable, show := shows.IsDownloadableShow(item.title)
			if isDownloadable {
				collectorConfig.Log.Debugf("Adding torrent %s", item.title)
				_, err = collectorConfig.TransmissionClient.AddTorrentByURL(item.Link, collectorConfig.DownloadDir)
				if err != nil {
					collectorConfig.Log.Errorf("Error adding torrent %s", err)
				}
				item.Show = show
				item.GetShowData()
				collectorConfig.Log.Debugf("Item %v", item)
			}
		}
		time.Sleep(5 * time.Minute)
	}

}
