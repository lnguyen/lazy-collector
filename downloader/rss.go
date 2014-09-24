package downloader

import (
	"regexp"
	"strconv"
	"time"

	"github.com/SlyMarbo/rss"
)

//RssFeedItem is an item in feed
type RssFeedItem struct {
	title     string
	Link      string
	Show      string
	Season    int
	Episode   int
	TimeAdded time.Time
}

//GetShowData determines Season/Episode
func (r *RssFeedItem) GetShowData() {
	re, _ := regexp.Compile(`S(\d+)E(\d+)`) // want to know what is in front of 'at'
	res := re.FindAllStringSubmatch(r.title, -1)
	r.Season, _ = strconv.Atoi(res[0][1])
	r.Episode, _ = strconv.Atoi(res[0][2])
}

//Feeds returns list of items in feed
func Feeds(feedURLs []string) ([]RssFeedItem, error) {
	var rssFeedItems []RssFeedItem
	for _, url := range feedURLs {
		feed, err := rss.Fetch(url)
		if err != nil {
			return []RssFeedItem{}, err
		}
		for _, item := range feed.Items {
			feedItem := RssFeedItem{title: item.Title, Link: item.Link}
			rssFeedItems = append(rssFeedItems, feedItem)
		}
	}
	return rssFeedItems, nil
}
