package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main(){
	file, _ := html.Parse(os.Stdin)
	fmt.Println(soleTitle(file))
}
func forEachNode(doc *html.Node, pre, post func(doc *html.Node)) {
	if pre != nil {
		pre(doc)
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(doc)
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
			fmt.Println("OK")
		case bailout{}:
			err = fmt.Errorf("несколько элементов title")
		default:
			panic(p)
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("нет элементов")
	}
	return title, nil
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path{
		if i > 0{
			sum += path[i-1].(path[i])
		}
	}
	return sum
}


func myFunction() (result int) {
    defer func() {
        if r := recover(); r != nil {
            result = r.(int) 
        }
    }()

    panic(42) 
}
