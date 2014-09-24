package downloader

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

type DownloadbleShowTest struct {
	shows    []string
	title    string
	expected bool
}

var downloadbleShowTest = []DownloadbleShowTest{
	{[]string{"south park", "game of thrones"}, "south park 123 foo bar caz", true},
	{[]string{"souTh paRk", "death of games"}, "sOuth death of", false},
	{[]string{"soUtH pArK", "game of thrones"}, "SouTh paRk 123 foo bar caz", true},
	{[]string{"soUtH pArK", "game of thrones"}, "Souh paRk 123 foo bar caz", false},
}

type MatchShowTest struct {
	show     string
	title    string
	expected bool
}

var matchShowTest = []MatchShowTest{
	{"south park", "south park 123 foo bar caz", true},
	{"south prk", "south park episode 242", false},
}

func TestIsDownloadableShow(t *testing.T) {
	for _, tt := range downloadbleShowTest {
		Convey("Test show is downloadable", t, func() {
			var shows Shows
			shows = tt.shows
			So(shows.IsDownloadableShow(tt.title), ShouldEqual, tt.expected)
		})
	}
}

func TestContains(t *testing.T) {
	Convey("Test if array contains element", t, func() {
		words := []string{"south", "root", "foo", "bar"}
		So(contains(words, "south"), ShouldEqual, true)
		So(contains(words, "root"), ShouldEqual, true)
		So(contains(words, "foo"), ShouldEqual, true)
		So(contains(words, "bar"), ShouldEqual, true)
		So(contains(words, "baz"), ShouldEqual, false)
	})
}

func TestMatchShow(t *testing.T) {
	for _, tt := range matchShowTest {
		Convey("Test shows are matched", t, func() {
			So(matchShow(tt.show, tt.title), ShouldEqual, tt.expected)
		})
	}
}
