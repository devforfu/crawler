package main

import (
    "../src/fetch"
    "flag"
    "golang.org/x/net/html"
    "log"
    "strings"
)

type NodeHandler = func (node *html.Node, depth int)

var url string
var maxDepth int

func main() {
    if strings.Trim(url, " ") == "" {
        log.Fatalf("The URL is empty")
    }
    result := fetch.Fetch(url)
    if result.Failed() {
        log.Fatalf("%s: %v", url, *result.Error)
    }
    doc, err := html.Parse(strings.NewReader(*result.Content))
    if err != nil {
        log.Fatalf("%s: %v", url, err)
    }
    forEachNode(doc, printElement, printElement, 0)
}

func init() {
    log.SetFlags(0)
    flag.StringVar(&url, "url", "", "URL to process")
    flag.IntVar(&maxDepth, "maxdepth", 10, "Search max. depth")
    flag.Parse()
}

func forEachNode(node *html.Node, pre, post NodeHandler, depth int) {
    if depth == maxDepth { return }
    if pre != nil { pre(node, depth) }
    for c := node.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post, depth + 1)
    }
    if post != nil { post(node, depth) }
}

func printElement(node *html.Node, depth int) {
    if node.Type == html.ElementNode {
        log.Printf("%*s<%s>\n", depth*2, "", node.Data)
    }
}
