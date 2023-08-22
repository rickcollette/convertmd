package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/russross/blackfriday/v2"
	"golang.org/x/net/html"
)

func main() {
	inputFile := flag.String("input", "", "Path to the input Markdown file")
	outputFile := flag.String("output", "", "Base path for the output files")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Please specify an input file using the -input flag.")
		return
	}

	if *outputFile == "" {
		fmt.Println("Please specify a base output path using the -output flag.")
		return
	}

	// Read the content of the Markdown file
	data, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	// Convert Markdown to HTML
	htmlData := blackfriday.Run(data)

	// Save as HTML
	err = os.WriteFile(*outputFile+".html", htmlData, 0644)
	if err != nil {
		fmt.Printf("Error writing HTML file: %s\n", err)
		return
	}

	// Extract plain text from HTML
	textData := htmlToText(string(htmlData))

	// Save as Text
	err = os.WriteFile(*outputFile+".txt", []byte(textData), 0644)
	if err != nil {
		fmt.Printf("Error writing text file: %s\n", err)
		return
	}

	fmt.Println("Conversion successful!")
}

func htmlToText(input string) string {
	doc, _ := html.Parse(strings.NewReader(input))
	var b strings.Builder
	htmlNodeToText(&b, doc)
	return b.String()
}

func htmlNodeToText(b *strings.Builder, n *html.Node) {
	if n.Type == html.TextNode {
		b.WriteString(n.Data)
	} else if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				b.WriteString(a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		htmlNodeToText(b, c)
	}
}
