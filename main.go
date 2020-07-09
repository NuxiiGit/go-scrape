package main

import "fmt"
import "io/ioutil"
import "log"
import "net/http"

func main() {
    response, err := http.Get("https://nuxiigit.github.io/README.txt")
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
