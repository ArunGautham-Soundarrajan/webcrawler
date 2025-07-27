package crawler

import (
	"fmt"
	stdio "io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ArunGautham-Soundarrajan/webcrawler/internal/io"
	"github.com/ArunGautham-Soundarrajan/webcrawler/internal/parser"

	"github.com/PuerkitoBio/goquery"
)

func Crawl(pageURL string, depth int, visited map[string]bool) {
	parsed, err := url.Parse(pageURL)
	if err != nil {
		fmt.Println("Invalid URL:", pageURL)
		return
	}

	if !CanCrawl(parsed.Path) {
		fmt.Println("Blocked by robots.txt:", pageURL)
		return
	}

	if visited[pageURL] || depth <= 0 {
		return
	}
	visited[pageURL] = true

	// Fetch the url
	resp, err := http.Get(pageURL)
	if err != nil {
		fmt.Println("Error fetching page:", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := stdio.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	bodyString := string(bodyBytes)

	// Create a Goquery document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(bodyString))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	title := doc.Find("title").Text()
	fmt.Println("URL:", pageURL, "Depth:", depth)

	domain, err := url.Parse(pageURL)
	if err != nil {
		fmt.Println("Error resolving url:", err)
		return
	}

	// Fetch Metadata
	metadata := GetMetadata(pageURL, resp, depth, title)

	content := parser.ExtractContentAsMarkDown(bodyString, domain.Hostname())

	// Save as MarkDown file
	filename := io.GenerateFilenameFromURL(pageURL)
	err = io.SavePageDataAsMarkdown(io.PageData{Metadata: metadata, Content: content}, filename+".md")
	if err != nil {
		fmt.Println("Error saving JSON:", err)
	}

	// Find the link and crawl recursively
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

func GetMetadata(pageURL string, resp *http.Response, depth int, title string) io.PageMetadata {
	return io.PageMetadata{
		URL:          pageURL,
		Title:        title,
		Date:         resp.Header.Get("Date"),
		LastModified: resp.Header.Get("Last-Modified"),
		CrawlTime:    time.Now().Format(time.RFC3339),
		Depth:        depth,
	}
}
