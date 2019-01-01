package main

import (
    "../src/links"
    "fmt"
    "log"
    "os"
)

func main() {
    links.BreadthFirst(crawl, os.Args[1:])
}

// crawl retrieves HTML page and collects links
func crawl(url string) []string {
    fmt.Println(url)
    list, err := links.Extract(url)
    if err != nil {
        log.Print(err)
    }
    return list
}
