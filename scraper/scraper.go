package scraper

import (
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
