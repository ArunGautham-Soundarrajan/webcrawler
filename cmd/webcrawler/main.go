package main

import (
	"net/url"

	"github.com/ArunGautham-Soundarrajan/webcrawler/internal/crawler"
)

func main() {
	var pageURL string = "https://www.scrapethissite.com/"
	var depth int = 2
	var visited = make(map[string]bool)

	parsed, _ := url.Parse(pageURL)
	crawler.RobotstxtInit(parsed.Scheme + "://" + parsed.Host)

	crawler.Crawl(pageURL, depth, visited)
}
