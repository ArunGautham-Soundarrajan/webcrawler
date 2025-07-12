package crawler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ArunGautham-Soundarrajan/webcrawler/internal/io"
	"github.com/ArunGautham-Soundarrajan/webcrawler/internal/parser"

	"github.com/PuerkitoBio/goquery"
)

func Crawl(pageURL string, depth int, visited map[string]bool) {
	if visited[pageURL] || depth <= 0 {
		return
	}
	visited[pageURL] = true

	resp, err := http.Get(pageURL)
	if err != nil {
		fmt.Println("Error fetching page:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	title := doc.Find("title").Text()
	fmt.Println("Title:", title, "URL:", pageURL, "Depth:", depth)

	content := parser.ExtractContent([]string{"h1", "h2", "p"}, doc)
	filename := io.GenerateFilenameFromURL(pageURL)
	err = io.SavePageDataAsJSON(io.PageData{URL: pageURL, Title: title, Content: content}, filename+".json")
	if err != nil {
		fmt.Println("Error saving JSON:", err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			refURL := ResolveURL(pageURL, href)
			if IsSameDomain(pageURL, refURL) {
				Crawl(refURL, depth-1, visited)
			}
		}
	})
}

func ResolveURL(base, ref string) string {
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}

	refURL, err := url.Parse(ref)
	if err != nil {
		return ""
	}

	return baseURL.ResolveReference(refURL).String()
}

func IsSameDomain(u1, u2 string) bool {
	host1, err1 := url.Parse(u1)
	host2, err2 := url.Parse(u2)

	if err1 != nil || err2 != nil {
		return false
	}

	return host1.Hostname() == host2.Hostname()
}