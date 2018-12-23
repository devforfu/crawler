package main

import (
    "./src/fetch"
    "fmt"
    "golang.org/x/net/html"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
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
        log.Printf("[%*d of %d] %s\n", width, i+1, total, url)
        err := WaitForServer(url)
        truncatedUrl := url[:minInt(30, len(url))]
        if err != nil {
            failed[url] = fmt.Sprintf("%s", err)
            continue
        }
        result := fetch.Fetch(url)
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
        log.Printf("%-30s | error: %s\n", url, strings.Trim(err, "\n\t\r"))
    }

    for url, count := range counts {
        log.Printf("%-30v | words: \t%d\timages:\t%d\n", url, count.words, count.images)
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

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
// Reference: gopl.io/ch5/wait
func WaitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil { return nil }
        log.Printf("server not responding (%s); retrying...", err)
        time.Sleep(time.Second << uint(tries))
    }
    return fmt.Errorf("server %s failed to responsd after %s", url, timeout)
}