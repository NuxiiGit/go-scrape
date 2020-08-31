package scraper

import (
    "os"
    "net/http"
    "io/ioutil"
)

func ReadURL(url string) ([]byte, error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    context, err := ioutil.ReadAll(response.Body)
    response.Body.Close()
    return context, err
}

func ReadFile(file string) ([]byte, error) {
    if _, err := os.Stat(file); os.IsNotExist(err) {
        return ReadURL(file)
    } else {
        return ioutil.ReadFile(file)
    }
}
