# RSS Feed Aggregator (Go)

A simple **RSS Feed Aggregator** built with **Go** that fetches, parses, and stores RSS feeds from multiple sources.  
It uses **concurrent scraping with goroutines**, a **PostgreSQL database** for persistence, and XML parsing to handle RSS data.

---

## Features

- Periodic scraping of RSS feeds using `time.Ticker`
- Concurrency with goroutines and `sync.WaitGroup`
- PostgreSQL database integration with `sqlc`
- Marks feeds as fetched after processing
- Parses RSS XML into Go structs
- Logs discovered feed items
- Error handling for network, parsing, and database operations



