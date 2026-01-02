package main

import (
    "context"
    "fmt"
    "golang.org/x/sync/errgroup"
    "net/http"
)

func fetchURL(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("failed to fetch %s: %w", url, err)
    }
    resp.Body.Close()
    return nil
}

func main() {
    urls := []string{"https://golang.org", "https://invalid.url", "https://example.com"}

    g, _ := errgroup.WithContext(context.Background())

    for _, url := range urls {
        url := url // capture variable
        g.Go(func() error {
            return fetchURL(url)
        })
    }

    err := g.Wait()
    if err != nil {
        fmt.Println("Error occurred:", err)
    } else {
        fmt.Println("Successfully fetched all URLs")
    }
}