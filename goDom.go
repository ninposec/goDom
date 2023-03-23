package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "regexp"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        input := scanner.Text()
        resp, err := http.Get(input)
        if err != nil {
            fmt.Fprintf(os.Stderr, "error making HTTP request to %s: %v\n", input, err)
            continue
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Fprintf(os.Stderr, "error reading response body: %v\n", err)
            continue
        }

        r := regexp.MustCompile(`(?i)https?://[\w.-]+\.\w+`)
        urls := r.FindAllString(string(body), -1)
        domains := make(map[string]bool)
        for _, url := range urls {
            subdomains := extractSubdomains(url)
            if len(subdomains) > 0 {
                domain := subdomains[len(subdomains)-1]
                if _, ok := domains[domain]; !ok {
                    domains[domain] = true
                    fmt.Println(url)
                }
            }
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}

func extractSubdomains(input string) []string {
    parsedURL, err := url.Parse(input)
    if err != nil {
        return nil
    }
    host := parsedURL.Hostname()
    var subdomains []string
    for i := len(host) - 1; i >= 0; i-- {
        if host[i] == '.' {
            subdomains = append(subdomains, host[:i])
        }
    }
    subdomains = append(subdomains, host)
    return subdomains
}
