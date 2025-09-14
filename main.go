package main

import (
	"fmt"
	"github.com/wrelin/web-crawler/internal/crawling"
	"github.com/wrelin/web-crawler/internal/report"
	"log"
	"os"
	"strconv"
)

const defaultMaxConcurrency = 3
const defaultMaxPages = 100

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]

	maxConcurrency := defaultMaxConcurrency
	if len(os.Args) >= 3 {
		converted, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error parse max concurrency to int: %v", err)
		}

		maxConcurrency = converted
	}

	maxPages := defaultMaxPages
	if len(os.Args) >= 4 {
		converted, err := strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatalf("Error parse max pages to int: %v", err)
		}

		maxPages = converted
	}

	cfg, err := crawling.Configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.Wg.Add(1)
	go cfg.CrawlPage(rawBaseURL)
	cfg.Wg.Wait()

	if err := report.WriteCSVReport(cfg.Pages, "report.csv"); err != nil {
		fmt.Printf("Error - writeCSVReport: %v\n", err)
	}
}
