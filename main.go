package main

import (
    "go-scrape/scraper"
    "os"
    "fmt"
    "bytes"
)

func main() {
    args := os.Args[1 :]
    argCount := len(args)
    if argCount < 1 {
        fmt.Println("usage:\n  go-scrape <url>")
    } else {
        var emitter string
        var url string
        if argCount == 1 {
            emitter = ""
            url = args[0]
        } else {
            emitter = args[0]
            url = args[1]
        }
        context, err := scraper.ReadFile(url)
        if err != nil {
            fmt.Printf("an error occurred when scraping:\n%s\n", err)
        } else if emitter == "" {
            fmt.Println(string(context))
        } else {
            node, err := scraper.DecodeHTML(context)
            var buffer bytes.Buffer
            if err != nil {
                fmt.Printf("an error occurred when decoding the page:\n%s\n", err)
            } else if emitter == "-json" {
                node.EncodeJSON(&buffer)
            }
            fmt.Printf("%s\n", &buffer)
        }
    }
}
