package report

import (
	"encoding/csv"
	"github.com/wrelin/web-crawler/internal/parsing"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func WriteCSVReport(pages map[string]parsing.PageData, filename string) error {
	columns := []string{
		"page_url",
		"h1",
		"first_paragraph",
		"outgoing_link_urls",
		"image_urls",
	}

	file, err := os.Create(filepath.Clean(filename))
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(columns)
	if err != nil {
		return err
	}

	// Sort keys for deterministic output
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		page := pages[key]
		row := []string{
			page.URL,
			page.H1,
			page.FirstParagraph,
			strings.Join(page.OutgoingLinks, ";"),
			strings.Join(page.ImageURLs, ","),
		}

		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
