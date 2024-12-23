package main

import (
	"fmt"
	"os"
	"strings"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findkeys: %v\n", err)
		os.Exit(1)
	}
	level := 0
	for _, node := range findNode(nil, doc, level) {
		if len(node) > 0 {
		fmt.Printf("[%s] ", node)
		}
	}
}

func findNode(node []string, doc *html.Node, level int) []string {
	// fmt.Println(doc.Data)
	if doc.Type == html.TextNode && doc.Parent.Data != "script" && doc.Parent.Data != "style" {
        text := strings.TrimSpace(doc.Data)
		node = append(node, text)
        // fmt.Println(text)
    }
	
	if doc.FirstChild != nil {
		node = findNode(node, doc.FirstChild, level)
	}
	if doc.NextSibling != nil {
		node = findNode(node, doc.NextSibling, level)
	}

	return node

}

