# web-crawler

Crawl site concurrently, save result to report.csv.

![Alt Text](demo.gif)

## Contributing

### Clone the repo
```bash
git clone https://github.com/wreling/web-crawler@latest
cd web-crawle
```

### Build the project
```bash
go mod tidy
go build
```

### Run the project
```bash
./web-crawler <site_url> <max_concurrent_thread> <max_pages_to_crawl>
```

### Run the tests
```bash
go test ./...
```