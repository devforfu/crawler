package main

import (
    "./src/fetch"
    "fmt"
    "golang.org/x/net/html"
    "os"
    "strings"
)

func main() {
    url := "https://gopl.io"
    result := fetch.Fetch(url)
    if result.Failed() { os.Exit(1) }
    words, images, err := CountWordsAndImages(*result.Content)
    if err != nil { os.Exit(1) }
    fmt.Printf("Number of words:\t%d\nNumber of images:\t%d", words, images)
}

func CountWordsAndImages(pageContent string) (words, images int, err error){
    doc, err := html.Parse(strings.NewReader(pageContent))
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc)
    return words, images, err
}

func countWordsAndImages(n *html.Node) (nWords, nImages int) {
    if n == nil { return }
    if n.Type == html.TextNode {
        text := n.Data
        words := strings.Split(strings.Trim(text, "\t\n\r"), " ")
        nWords += len(words)
    } else if n.Type == html.ElementNode && n.Data == "img" {
        nImages += 1
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        w, i := countWordsAndImages(c)
        nWords += w
        nImages += i
    }
    return nWords, nImages
}