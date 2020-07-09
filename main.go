package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
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
    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    context, err := ioutil.ReadAll(response.Body)
    response.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", context)
}
