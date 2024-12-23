package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, tags ...string) []*html.Node {
	var result []*html.Node
	element := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, tag := range tags {

				if n.Data == tag {
					result = append(result, n)
					break
				}

			}
		}
	}
	
	findTag(doc, element)
	return result

}

func findTag(doc *html.Node, elem func(doc *html.Node)) {
	if elem != nil {
		elem(doc)
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		findTag(c, elem)
	}
}

func main() {
	doc := `<html>
    <head>
        <title>Example</title>
    </head>
    <body>
        <h1>$title</h1>
        <p>Paragraph</p>
    </body>
</html>`
	reader := strings.NewReader(doc)
	d, _ := html.Parse(reader)

	fmt.Println(len(ElementsByTagName(d, "h1", "p", "title")))

}
