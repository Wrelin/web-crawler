package crawling

import (
	"fmt"
	"net/url"

	"github.com/wrelin/web-crawler/internal/parsing"
)

func (cfg *Config) CrawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.Wg.Done()
	}()

	if cfg.pagesLen() >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - CrawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	// stay within the same site
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}

	// Only proceed the first time we see this normalized URL
	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	// Extract all the data we care about and store it
	pageData := parsing.ExtractPageData(htmlBody, rawCurrentURL)
	cfg.setPageData(normalizedURL, pageData)

	// Recurse using the already-extracted outgoing links
	for _, nextURL := range pageData.OutgoingLinks {
		cfg.Wg.Add(1)
		go cfg.CrawlPage(nextURL)
	}
}
