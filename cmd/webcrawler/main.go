package main

import (
	"github.com/ArunGautham-Soundarrajan/webcrawler/internal/crawler"
)

func main() {
	var pageURL string = "https://www.scrapethissite.com/"
	var depth int = 2
	var visited = make(map[string]bool)

	crawler.Crawl(pageURL, depth, visited)
}