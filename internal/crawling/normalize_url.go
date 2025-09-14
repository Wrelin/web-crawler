package crawling

import (
	"net/url"
	"strings"
)

func normalizeURL(rawUrl string) (string, error) {
	parsed, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	normalized := parsed.Host + parsed.Path
	normalized = strings.ToLower(normalized)

	return strings.TrimSuffix(normalized, "/"), nil
}
