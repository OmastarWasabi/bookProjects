package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Рекурсивная функция для отображения всех узлов
func printNode(n *html.Node, level int) {
	// Печать текущего узла
	fmt.Printf("%*sNode: %+v\n", level*2, "", n)

	// Если есть дочерний узел, рекурсивно обрабатываем его
	if n.FirstChild != nil {
		printNode(n.FirstChild, level+1)
	}
	// Если есть следующий узел на том же уровне, также обрабатываем его
	if n.NextSibling != nil {
		printNode(n.NextSibling, level)
	}
	fmt.Printf("LEVEL: %d", level)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	printNode(doc, 0)
}