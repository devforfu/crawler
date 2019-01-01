package main

import (
    "../src/fetch"
    "fmt"
    "os"
)

func main() {
    urls := os.Args[1:]
    if len(urls) < 1 {
        fmt.Fprintf(os.Stderr, "fetch: no URLs to fetch\n")
        os.Exit(1)
    }
    for _, url := range urls {
        result := fetch.Fetch(url)
        if result.Success {
            fmt.Fprintf(os.Stdout, *result.Content)
        } else {
            fmt.Fprintf(os.Stderr, *result.Error)
        }
    }
}