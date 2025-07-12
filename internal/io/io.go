package io

import (
	"encoding/json"
	"net/url"
	"os"
	"strings"
)

type PageData struct {
	URL     string              `json:"url"`
	Title   string              `json:"title"`
	Content map[string][]string `json:"content"`
}

func SavePageDataAsJSON(pageData PageData, filename string) error {
	err := os.MkdirAll("output", 0755)
	if err != nil {
		return err
	}

	file, err := os.Create("output/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(pageData)
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