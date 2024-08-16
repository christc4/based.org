package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"os"
)

// fetchHTML fetches the HTML content of a given URL
func fetchHTML(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// extractTextFromPTags recursively extracts text from <p> tags in the HTML node
func extractTextFromPTags(n *html.Node) {
	// If the node is a <p> tag
	if n.Type == html.ElementNode && n.Data == "p" {
		// Extract and print the text inside the <p> tag
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				fmt.Println(c.Data)
			}
		}
	}

	// Recursively process all child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractTextFromPTags(c)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ptag_scraper <URL>")
		return
	}

	url := os.Args[1]

	doc, err := fetchHTML(url)
	if err != nil {
		fmt.Println("Error fetching HTML:", err)
		return
	}

	extractTextFromPTags(doc)
}

