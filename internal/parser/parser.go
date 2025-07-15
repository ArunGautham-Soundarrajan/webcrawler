package parser

import (
	"strings"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/PuerkitoBio/goquery"
)

func ExtractContent(tags []string, doc *goquery.Document) map[string][]string {
	content := make(map[string][]string)
	for _, tag := range tags {
		doc.Find(tag).Each(func(i int, s *goquery.Selection) {
			text := strings.TrimSpace(s.Text())
			if text != "" {
				content[tag] = append(content[tag], text)
			}
		})
	}
	return content
}

func ExtractContentAsMarkDown(body string, domain string) string {

	markdown, err := htmltomarkdown.ConvertString(
		body,
		converter.WithDomain(domain),
	)

	if err != nil {
		return ""
	}

	return markdown

}
