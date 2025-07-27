package main

import (
	"net/url"

	"github.com/ArunGautham-Soundarrajan/webcrawler/internal/crawler"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "webcrawler",
	Short: "A simple web crawler",
	Run: func(cmd *cobra.Command, args []string) {
		PageUrl, _ := cmd.Flags().GetString("url")
		depth, _ := cmd.Flags().GetInt("depth")
		start(PageUrl, depth)
	},
}

func init() {
	rootCmd.Flags().String("url", "", "URL to crawl")
	rootCmd.Flags().Int("depth", 1, "Depth to crawl")
	rootCmd.MarkFlagRequired("url")
}
func start(PageUrl string, depth int) {

	parsed, _ := url.Parse(PageUrl)
	crawler.RobotstxtInit(parsed.Scheme + "://" + parsed.Host)
	var visited = make(map[string]bool)

	crawler.Crawl(PageUrl, depth, visited)

}

func main() {
	rootCmd.Execute()
}
