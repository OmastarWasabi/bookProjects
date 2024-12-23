package main

import (
	"strings"

	"golang.org/x/net/html"
)

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
    ElementById(d, "h1")	
	
}

// func expand(s string, f func(string)string) string{
// 	var result strings.Builder
// 	i:=0
// 	for i < len(s){
// 		j := i+1
// 		if s[i] == '$'{
			
// 			for j < len(s) && (s[j] > 'a' && s[j] < 'z' || s[j] > 'A' && s[j] < 'Z'){
// 				j++
// 			} 
// 		}
// 		if j > i+1{
// 			key := s[i+1:j]
// 			result.WriteString(f(key))
// 			i = j
// 			continue
// 		}
// 		result.WriteByte(s[i])
// 		i++
// 	}
// 	return result.String()
// }

func ElementById(doc *html.Node, id string) *html.Node{
	var result *html.Node
	startElement := func (n *html.Node) bool{
		if n.Type == html.ElementNode{
			if n.Data == id{
				result = n
				return false
			}
		}
		return true
	}
	forEachNode(doc, startElement, nil)
	return result

}

func forEachNode(doc *html.Node, pre, post func(doc *html.Node) bool) {
	if pre != nil && !pre(doc){
		return
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil && !post(doc) {
		return
	}
}

// var level int

// func startElement(n *html.Node) {
// 	switch n.Type {
// 	case html.ElementNode:
// 		attrs := []string{}
// 		for _, attr := range n.Attr {
// 			attrs = append(attrs, fmt.Sprintf("%s=%q", attr.Key, attr.Val))
// 		}
// 		attrStr := ""
// 		if len(attrs) > 0 {
// 			attrStr = " " + strings.Join(attrs, " ")
// 		}

// 		if n.FirstChild == nil {
// 			fmt.Printf("%*s<%s%s/>\n", level*2, "", n.Data, attrStr)
// 		} else {
// 			fmt.Printf("%*s<%s%s/>\n", level*2, "", n.Data, attrStr)
// 			level++
// 		}
// 	case html.TextNode:
// 		text := strings.TrimSpace(n.Data)
// 		if text != ""{
// 			fmt.Printf("%s%s", level*2, "", text)
// 		}
// 	case html.CommentNode:
// 		fmt.Printf("%*s<<-- %s -->>\n", level*2, "", n.Data)	
// 	}

// }

// func endElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		level--
// 		fmt.Printf("%*s</%s>\n", level*2, "", n.Data)
// 	}
// }
