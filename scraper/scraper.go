// Provides functions for scraping and parsing local and remote files.
package scraper

import (
    "os"
    "net/http"
    "io/ioutil"
)

// Attempts to scrape a remote file if this URL points to an accessible webpage.
func ReadURL(url string) ([]byte, error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    context, err := ioutil.ReadAll(response.Body)
    response.Body.Close()
    return context, err
}

// Attempts to scrape a local file if it exists, otherwise it assumes the filepath is a URL.
func ReadFile(file string) ([]byte, error) {
    if _, err := os.Stat(file); err == nil {
        return ioutil.ReadFile(file)
    } else {
        return ReadURL(file)
    }
}
