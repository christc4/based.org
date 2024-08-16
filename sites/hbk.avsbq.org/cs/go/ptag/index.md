# Parse html for p tag

last update 2024/08/15

Source code availabe, [main.go](main.go)

I will breakdown what this small program does and explain some lines
I most likely will massively improve this so it gets information from all HTML tags and prints it in a much neater way...

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

	extractTextFromPTags(doc)
	}


## Stuff explained

### packages

`net/http` provides HTTP client and server implementations, used here to make HTTP requests

`golang.org/x/net/html` provide function to parse HTML documents

`func fetchHTML(url string) (*html.Node, error) {`

> defines new function, `fetchHTML`, that takes a single argument `url` of the type `string`, returning 2 values: a pointer to `html.Node` and an `error`

`resp, err := http.Get(url)`

> makes a HTTP GET request to the specified `url` using the `http.Get` function, `resp` is a variable storing the HTTP response returned by `http.Get`, `err` a variable storing any error encountered during the GET request, `:=` short variable declaration syntax in Go that declares and initializes `resp` and `err` in a single step

	if err != nil {
		return nil, err
	}


> block that checks if there are errors during HTTP request

- `if err != nil` condition that checks if `err` is not `nil` (i.e., if there was an error)
- `return nil, err` if error has occured, function returns `nil` (no HTML document) and the error itself

`defer resp.Body.Close()`

> Ensures that the HTTP response body is closed after we're done using it, evin if errors occur later

- `defer` this keyword schedules the `resp.Body.Close()` function to run after the `fetchHTML` function completes, ensuring that resources are properly released

`doc, err := html.Parse(resp.Body)`

> parses the HTML document from the response body

- `html.Parse(resp.body)` function that parses the HTML from the response body returning a pointer to `html.Node` representing the root of the document
- `doc` variable that stores the parsed HTMl document

	if err != nil {
		return nil, err
	}

> this block checks if there was an error during the HTML parsing

- if an error occurs, the function returns `nil` and the error

`return doc, nil`

> this line returns the parsed HTML document `doc` and `nil` for the error

`func extractTextFromPTags(n *html.Node) {`

> this line recursively extracts text from p tags in the `html.Node`, the new function is defined taking a single argument `n` which is a pointer to `html.Node`, this function tarverses the HTML tree extracting text from p tags

`if n.Type == html.ElementNode && n.Date == "p" {`

> this line checks if the current node is an HTML element and if that element is a p tag

- `n.Type == html.ElementNode` checks if the node type is an element (as opposed to text, comment, etc.)

- `n.Data == "p"` checks if the element is specifically a p tag

`for c := n.FirstChild; c != nil; c = c.NextSibling {`

> This line starts a `for` loop to iterate over the child nodes of the p tag.

- `c := n.FirstChild` initializes `c` to the first child of the current node
- `c != nil` the loop continues as long as `c` is not `nil`
- `c = c.NextSibling` moves `c` to the next sibling node after each iteration

`if c.Type == html.TextNode { fmt.Println(c.Data) }`

> this block checks if the child node is a text node and prints its data

`for c := n.FirstChild; c != nil; c = c.NextSibling { extractTextFromPTags(c) }`

> This block is another `for` loop that recursively calls `extractTextFromPTags` on each child node. This allows the function to traverse the entire HTML document tree, extracting text from all p tags, no matter  how deeply nested they are

`func main() {`

> this line defines the `main` function, which is the entry point of the Go program. When the program is executed the `main` function isi called first

` if len(os.Args) < 2 { fmt.Println("Usage: ptag_scraper <URL>") return } `

> this block checks if a URL argument was provided when the program has ran

- `len(os.Args) < 2` checks if fewer than 2 command line arguments were passed (the first argument is always the program name such as ./ptag)
- `return` exits the `main` function early if no URL is provided

`url := os.Args[1]`

- 

## Basics

`mkdir ptag`

`cd ptag`

`go mod init ptag`

`go get golang.org/x/net/html`

`go build -o ptag_scraper`

`./ptag https://avsbq.org` ;)
