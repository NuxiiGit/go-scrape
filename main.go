package main

import (
    "go-scrape/scraper"
    "os"
    "log"
    "fmt"
)

func main() {
    args := os.Args[1:]
    var url string
    if len(args) < 1 {
        log.Print("please supply a url to scrape")
        url = "https://nuxiigit.github.io/README.txt"
    } else {
        url = args[0]
    }
    context, err := scraper.ReadURL(url)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", context)
}
