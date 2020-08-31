package main

import (
    "go-scrape/scraper"
    "os"
    "fmt"
)

func main() {
    args := os.Args[1 :]
    if len(args) < 1 {
        fmt.Println("usage:\n  go-scrape <url>")
    } else {
        url := args[0]
        context, err := scraper.ReadFile(url)
        if err != nil {
            fmt.Printf("an error occurred when scraping:\n%s\n", err)
        } else {
            fmt.Printf("%s", context)
        }
    }
}
