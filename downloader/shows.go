package downloader

import "strings"

//Shows is list of shows
type Shows []string

//IsDownloadableShow determine if show is downloadable based on title
func (s Shows) IsDownloadableShow(title string) (bool, string) {
	var isMatched bool
	var returnedShow string
	title = strings.ToLower(title)
	for _, show := range s {
		show = strings.ToLower(show)
		isMatched = matchShow(show, title)
		if isMatched {
			returnedShow = show
			break
		}
	}
	return isMatched, returnedShow
}

func matchShow(show string, title string) bool {
	tokenizedShow := strings.Split(show, " ")
	tokenizeTitle := strings.Split(title, " ")
	match := 0
	for _, word := range tokenizedShow {
		if contains(tokenizeTitle, word) {
			match++
		}
	}
	if match == len(tokenizedShow) {
		return true
	}
	return false
}

func contains(array []string, element string) bool {
	for _, e := range array {
		if e == element {
			return true
		}
	}
	return false
}
