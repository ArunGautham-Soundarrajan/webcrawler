# Web Crawler

A web crawler written in Go that extracts content from web pages and saves it as Markdown files.

## Features

- Crawls web pages recursively
- Extracts content from HTML pages
- Saves extracted content as Markdown files
- Includes metadata (title, URL, date, last modified, crawl time, and depth) in YAML frontmatter

## Example Output

The crawler saves extracted content as Markdown files with YAML frontmatter. Here's an example of what the output might look like:

```markdown
---
title: Example Web Page
url: https://www.example.com
date: 2022-01-01
lastModified: 2022-01-01
crawlTime: 2022-01-01T12:00:00Z
depth: 1
---

# Example Web Page

This is an example web page.

## Section 1

This is the first section.

## Section 2

This is the second section.
```

## Todo

- [x] Add support for robots.txt
- [ ] Add support for multiple crawl threads
- [ ] Implement rate limiting to avoid overwhelming websites
- [ ] Add option to save output to a database instead of filesystem
