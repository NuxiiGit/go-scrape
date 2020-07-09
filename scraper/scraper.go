package scraper

import (
    "net/http"
    "io/ioutil"
)

func ReadURL(url string) (context []byte, err error) {
    response, err := http.Get(url)
    defer response.Body.Close()
    if err != nil {
        return nil, err
    }
    return ioutil.ReadAll(response.Body)
}
