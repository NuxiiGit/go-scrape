package scraper

import (
    "os"
    "net/http"
    "io/ioutil"
)

func ReadURL(url string) (context []byte, err error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    a, b := ioutil.ReadAll(response.Body)
    response.Body.Close()
    return a, b
}

func ReadFile(file string) (context []byte, err error) {
    if _, err := os.Stat(file); os.IsNotExist(err) {
        return ReadURL(file)
    } else {
        return ioutil.ReadFile(file)
    }
}
