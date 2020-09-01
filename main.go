package main

import (
    "go-scrape/scraper"
    "os"
    "fmt"
    "bytes"
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
            node, err := scraper.DecodeHTML(context)
            if err != nil {
                fmt.Printf("an error occurred when decoding the page:\n%s\n", err)
            } else {
                var buffer bytes.Buffer
                node.EncodeJSON(&buffer)
                fmt.Printf("%s\n", &buffer)
                fmt.Printf("%s\n", &node)
            }
        }
    }
}
