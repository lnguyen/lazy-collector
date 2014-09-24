package downloader

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

//Blacklist is list of shows already downloaded
type Blacklist []RssFeedItem

//AddToBlackList add rss item to feed
func AddToBlackList(item RssFeedItem) {
	blacklist := getCurrentBlacklist()
	blacklist = append(blacklist, item)
	writeBlackList(blacklist)
}

//Blacklisted see if item is blacklisted
func Blacklisted(rssFeedItem RssFeedItem) bool {
	blacklist := getCurrentBlacklist()
	for _, item := range blacklist {
		if (rssFeedItem.Show == item.Show) &&
			(rssFeedItem.Season == item.Season) &&
			(rssFeedItem.Episode == item.Episode) {
			return true
		}
	}
	return false
}

func writeBlackList(blacklist Blacklist) {
	os.Mkdir(filepath.Join(getHomeDir(), ".lazy"), 0755)
	file, err := os.Create(blacklistFile())
	defer file.Close()
	if err != nil {
		Log.Debugf("Error creating blacklist file %s", err)
	}
	blacklistMarshalled, err := json.Marshal(blacklist)
	if err != nil {
		Log.Debugf("Error marshaling empty blacklist %s", err)
	}
	file.Write(blacklistMarshalled)
}

func getCurrentBlacklist() Blacklist {
	var blackList Blacklist
	if _, err := os.Stat(blacklistFile()); err != nil {
		var emptyList Blacklist
		emptyList = []RssFeedItem{}
		writeBlackList(emptyList)
	}

	blackListByte, err := ioutil.ReadFile(blacklistFile())
	if err != nil {
		Log.Debug("Unable to read blacklist file")
	}
	err = json.Unmarshal(blackListByte, &blackList)
	if err != nil {
		Log.Debug("Unable to unmarshal")
	}
	return blackList
}

func blacklistFile() string {
	return filepath.Join(getHomeDir(), ".lazy/blacklist.json")
}

func getHomeDir() string {
	return os.Getenv("HOME")
}
