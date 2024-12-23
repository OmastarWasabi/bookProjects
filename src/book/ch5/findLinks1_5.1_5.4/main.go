package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	// если элемент является ссылкой <a>, добавляем href в список ссылок
	if n.Type == html.ElementNode {
		//добавляем к поиску ссылок теги  img, script, link
		switch n.Data {
		case "a":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		case "img", "script":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		case "link":
			var flag bool
			href := ""
			for _, link := range n.Attr {
				if link.Key == "rel" && link.Val == "stylesheet" {
					flag = true
				}
				if link.Key == "href" {
					href = link.Val

				}
			}
			if flag && href != ""{
				links = append(links, href)
			}
		}
	}

	// рекурсивно обходим дочерние узлы через FirstChild
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	// после обхода всех дочерних узлов продолжаем на том же уровне через NextSibling
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	fmt.Println(n.Data)

	// Обход всех дочерних узлов с помощью цикла
	// for c := n.FirstChild; c != nil; c = c.NextSibling{
	// 	links = visit(links, c)
	// }

	return links
}


// Дерево документа issuehtml
// <h1>34990 тем</h1>
// └── Текст: "34990 тем"

// <table>
// ├── <tr style='text-align: left'>
// │   ├── <th>#</th>
// │   │   └── Текст: "#"
// │   ├── <th>State</th>
// │   │   └── Текст: "State"
// │   ├── <th>User</th>
// │   │   └── Текст: "User"
// │   ├── <th>Titlr</th>
// │   │   └── Текст: "Titlr"
// ├── <tr>
// │   ├── <td>
// │   │   └── <a href='https://github.com/HUIIIM/Auto-GPT/pull/14'>14</a>
// │   │       └── Текст: "14"
// │   ├── <td>open</td>
// │   │   └── Текст: "open"
// │   ├── <td>
// │   │   └── <a href='https://github.com/HUIIIM'>HUIIIM</a>
// │   │       └── Текст: "HUIIIM"
// │   ├── <td>
// │   │   └── <a href='https://github.com/HUIIIM/Auto-GPT/pull/14'>hui</a>
// │   │       └── Текст: "hui"
// ├── <tr>
// │   ├── <td>
// │   │   └── <a href='https://github.com/Liana10042024/team_project_mini_ai/pull/1'>1</a>
// │   │       └── Текст: "1"
// │   ├── <td>open</td>
// │   │   └── Текст: "open"
// │   ├── <td>
// │   │   └── <a href='https://github.com/eunhui33'>eunhui33</a>
// │   │       └── Текст: "eunhui33"
// │   ├── <td>
// │   │   └── <a href='https://github.com/Liana10042024/team_project_mini_ai/pull/1'>Hui</a>
// │   │       └── Текст: "Hui"
