package main

import (
    "./src/fetch"
    "fmt"
    "golang.org/x/net/html"
    "os"
    "strconv"
    "strings"
)

type Counts = struct { words, images int }

func main() {
    fmt.Printf("Counting words and images in URLs\n")
    urls := os.Args[1:]
    total := len(urls)
    width := len(strconv.Itoa(total))
    failed := map[string]string {}
    counts := map[string]Counts {}

    for i, url := range urls {
        fmt.Printf("[%*d of %d] %s\n", width, i+1, total, url)
        result := fetch.Fetch(url)
        truncatedUrl := url[:minInt(30, len(url))]
        if result.Failed() {
            failed[truncatedUrl] = *result.Error
            continue
        }
        words, images, err := CountWordsAndImages(*result.Content)
        if err != nil {
            failed[truncatedUrl] = fmt.Sprintf("%s", err)
            continue
        }
        counts[truncatedUrl] = Counts{words, images}
    }

    for url, err := range failed {
        fmt.Printf("%-30s | error: %s\n", url, strings.Trim(err, "\n\t\r"))
    }

    for url, count := range counts {
        fmt.Printf("%-30v | words: \t%d\timages:\t%d\n", url, count.words, count.images)
    }
}

func minInt(a, b int) int {
    if a < b { return a }
    return b
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