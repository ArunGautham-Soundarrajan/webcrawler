package crawler

import (
	"net/http"

	"github.com/temoto/robotstxt"
)

var robotsData *robotstxt.RobotsData

func RobotstxtInit(domain string) {
	resp, err := http.Get(domain + "/robots.txt")
	if err != nil {
		robotsData = nil
		return
	}
	defer resp.Body.Close()

	data, err := robotstxt.FromResponse(resp)
	if err != nil {
		robotsData = nil
		return
	}
	robotsData = data
}

func CanCrawl(url string) bool {
	if robotsData == nil {
		return true
	}
	group := robotsData.FindGroup("*")
	return group.Test(url)
}
