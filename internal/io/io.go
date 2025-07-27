package io

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type PageMetadata struct {
	URL          string `yaml:"url"`
	Title        string `yaml:"title"`
	Date         string `yaml:"date,omitempty"`
	LastModified string `yaml:"last_modified,omitempty"`
	CrawlTime    string `yaml:"crawl_time"`
	Depth        int    `yaml:"depth"`
}

type PageData struct {
	Metadata PageMetadata `json:"metadata"`
	Content  string       `json:"content"`
}

func SavePageDataAsMarkdown(pageData PageData, filename string) error {
	err := os.MkdirAll("output", 0755)
	if err != nil {
		return err
	}

	file, err := os.Create("output/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	metaYAML, err := yaml.Marshal(pageData.Metadata)
	if err != nil {
		return err
	}

	// Compose frontmatter with YAML delimiters
	frontmatter := fmt.Sprintf("---\n%s---\n\n", string(metaYAML))

	_, err = file.WriteString(frontmatter + pageData.Content)
	return err

}

func GenerateFilenameFromURL(rawurl string) string {
	parsed, err := url.Parse(rawurl)
	if err != nil {
		return "page"
	}
	name := strings.ReplaceAll(parsed.Path, "/", "_")
	if name == "" || name == "_" {
		name = "index"
	}
	return parsed.Hostname() + name
}
