package main

import (
    "go-scrape/scraper"
    "go-scrape/data"
    "os"
    "fmt"
    "bytes"
)

func printHelp() {
    fmt.Println("usage:\n  go-scrape <url> [json|xml]")
}

func main() {
    args := os.Args[1 :]
    argCount := len(args)
    if argCount < 1 {
        printHelp()
    } else {
        url := args[0]
        context, err := scraper.ReadFile(url)
        if err != nil {
            fmt.Printf("an error occurred when scraping!\n%s\n", err)
        } else if argCount == 1 {
            fmt.Println(string(context))
        } else {
            emitter := args[1]
            node, err := data.DecodeHTML(context)
            if err != nil {
                fmt.Printf("an error occurred when decoding the page!\n%s\n", err)
            } else if emitter == "json" {
                var buffer bytes.Buffer
                node.EncodeJSON(&buffer)
                fmt.Print(&buffer)
            } else if emitter == "xml" {
                var buffer bytes.Buffer
                node.EncodeXML(&buffer)
                fmt.Print(&buffer)
            } else {
                fmt.Printf("unknown renderer '%s'\n", emitter)
                printHelp()
            }
        }
    }
}
