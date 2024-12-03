package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func fetchOGP(url string) (siteName string, title string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch OGP URL: %w", err)
	}
	defer resp.Body.Close()

	//body, err := io.ReadAll(resp.Body)
	//log.Println(string(body))

	limitedReader := io.LimitReader(resp.Body, 3072)
	doc, err := html.Parse(limitedReader)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	var findOGP func(*html.Node)
	findOGP = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			var property, content string
			for _, attr := range n.Attr {
				if attr.Key == "property" {
					property = attr.Val
				} else if attr.Key == "content" {
					content = attr.Val
				}
			}
			switch property {
			case "og:site_name":
				siteName = content
			case "og:title":
				title = content
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findOGP(c)
		}
	}

	findOGP(doc)

	if siteName == "" && title == "" {
		return "", "",
			fmt.Errorf("OGP data none")
	}
	return siteName, title, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go https://soulminingrig.com/")
		return
	}

	url := os.Args[1]
	siteName, title, err := fetchOGP(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Site Name: %s\n", siteName)
	fmt.Printf("Title: %s\n", title)
}
