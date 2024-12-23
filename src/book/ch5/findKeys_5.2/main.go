package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main(){
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findkeys: %v\n", err)
		os.Exit(1)
	}
	tags := make(map[string]int)
	fmt.Println(findKeys(tags, doc))

}

func findKeys(keys map[string]int,doc *html.Node) map[string]int {
	 
	if doc.Type == html.ElementNode {
		keys[doc.Data]++
	}
	if doc.FirstChild != nil {
		keys = findKeys(keys, doc.FirstChild)
	}
	if doc.NextSibling != nil {
		keys = findKeys(keys, doc.NextSibling)
	}
	return keys
}