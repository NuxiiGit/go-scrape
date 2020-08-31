package main

import (
    "go-scrape/scraper"
    "os"
    "fmt"
)

func main() {
    args := os.Args[1 :]
    if len(args) < 1 {
        fmt.Print("usage:\n  go-scrape <url>\n")
    } else {
        url := args[0]
        context, err := scraper.ReadURL(url)
        if err != nil {
            fmt.Printf("an error occurred when reading the url -- %s", err)
        } else {
            fmt.Printf("%s", context)
        }
    }
}
