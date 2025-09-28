package crawling

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/wrelin/web-crawler/internal/parsing"
)

type Config struct {
	Pages              map[string]parsing.PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
	maxPages           int
}

// addPageVisit returns true if this is the first time we see the URL.
// We insert a placeholder PageData to mark it as visited; it’ll be replaced later.
func (cfg *Config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.Pages[normalizedURL]; visited {
		return false
	}

	cfg.Pages[normalizedURL] = parsing.PageData{URL: normalizedURL}
	return true
}

// setPageData safely stores the final PageData for a URL.
func (cfg *Config) setPageData(normalizedURL string, data parsing.PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.Pages[normalizedURL] = data
}

func Configure(rawBaseURL string, maxConcurrency, maxPages int) (*Config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &Config{
		Pages:              make(map[string]parsing.PageData),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		Wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil
}

func (cfg *Config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.Pages)
}
