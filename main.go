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
        fmt.Println("usage:\n  go-scrape <url> [json|xml]")
    } else {
        url := args[0]
        context, err := scraper.ReadFile(url)
        if err != nil {
            fmt.Printf("an error occurred when scraping:\n%s\n", err)
        } else if argCount == 1 {
            fmt.Println(string(context))
        } else {
            emitter := args[1]
            node, err := scraper.DecodeHTML(context)
            var buffer bytes.Buffer
            if err != nil {
                fmt.Printf("an error occurred when decoding the page:\n%s\n", err)
            } else if emitter == "json" {
                node.EncodeJSON(&buffer)
            } else if emitter == "xml" {
                node.EncodeXML(&buffer)
            } else {
                fmt.Printf("unknown emitter: %s\n", emitter)
            }
            fmt.Printf("%s\n", &buffer)
        }
    }
}
