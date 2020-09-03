// Provides functions for scraping and parsing local and remote files.
package scraper

import (
    "os"
    "net/http"
    "io/ioutil"
    "bytes"
    "strings"
    "encoding/xml"
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

// Represents a HTML element.
type HTMLElement struct {
    XMLName xml.Name
    Attrs []xml.Attr `xml:",any,attr"`
    Content []byte `xml:",chardata"`
    Children []HTMLElement `xml:",any"`
}

func (elem *HTMLElement) UnmashalXML(decoder *xml.Decoder, start xml.StartElement) error {
    elem.Attrs = start.Attr
    type htmlelement HTMLElement
    return decoder.DecodeElement((*htmlelement)(elem), &start)
}

func sanitiseString(dirty string) string {
    phase1 := strings.Replace(dirty, "\\", "\\\\", -1)
    phase2 := strings.Replace(phase1, "\"", "\\\"", -1)
    phase3 := strings.Replace(phase2, "\n", "\\n", -1)
    phase4 := strings.Replace(phase3, "\r", "\\r", -1)
    return phase4
}

func sanitiseInnerHTML(dirty string) string {
    phase1 := strings.Replace(dirty, "\n", "&#xD;", -1)
    phase2 := strings.Replace(phase1, "\r", "&#xA;", -1)
    return phase2
}

// Encodes a HTML element into XML.
func (elem *HTMLElement) EncodeXML(buffer *bytes.Buffer) {
    buffer.WriteString(`<`)
    buffer.WriteString(sanitiseInnerHTML(elem.XMLName.Local))
    for _, attr := range elem.Attrs {
        buffer.WriteString(` `)
        buffer.WriteString(sanitiseInnerHTML(attr.Name.Local))
        buffer.WriteString(`="`)
        buffer.WriteString(sanitiseString(attr.Value))
        buffer.WriteString(`"`)
    }
    buffer.WriteString(`>`)
    buffer.WriteString(sanitiseInnerHTML(string(elem.Content)))
    buffer.WriteString(`</`)
    buffer.WriteString(sanitiseInnerHTML(elem.XMLName.Local))
    buffer.WriteString(`>`)
}

// Encodes a HTML element into JSON.
func (elem *HTMLElement) EncodeJSON(buffer *bytes.Buffer) {
    buffer.WriteString(`{"name":"`)
    buffer.WriteString(sanitiseString(elem.XMLName.Local))
    buffer.WriteString(`","attrs":[`)
    for _, attr := range elem.Attrs {
        buffer.WriteString(`{"name":"`)
        buffer.WriteString(sanitiseString(attr.Name.Local))
        buffer.WriteString(`","value":"`)
        buffer.WriteString(sanitiseString(attr.Value))
        buffer.WriteString(`"},`)
    }
    buffer.WriteString(`],"content":"`)
    buffer.WriteString(sanitiseString(string(elem.Content)))
    buffer.WriteString(`","children":[`)
    for _, child := range elem.Children {
        child.EncodeJSON(buffer)
        buffer.WriteString(",")
    }
    buffer.WriteString(`],`)
    buffer.WriteString(`}`)
}

// Decodes a HTML page into a tree structure.
func DecodeHTML(html []byte) (HTMLElement, error) {
    buffer := bytes.NewBufferString(`<root>`)
    buffer.Write(html)
    buffer.WriteString(`</root>`)
    decoder := xml.NewDecoder(buffer)
    decoder.Strict = false
    decoder.AutoClose = xml.HTMLAutoClose
    decoder.Entity = xml.HTMLEntity
    var root HTMLElement
    err := decoder.Decode(&root)
    if err != nil {
        return root, err
    }
    return root, nil
}
