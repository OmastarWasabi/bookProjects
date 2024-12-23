package example

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)


func main(){
	for _, url := range os.Args[1:]{
		links, err := findlinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "finlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
func findlinks(url string) ([]string, error){
	resp, err := http.Get(url)
	if err != nil{
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("получение %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("анализ %s как HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	// если элемент является ссылкой <a>, добавляем href в список ссылок
	if n.Type == html.ElementNode {
		// добавляем к поиску ссылок теги  img, script, link
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

	// обход всех дочерних узлов с помощью цикла
	// for c := n.FirstChild; c != nil; c = c.NextSibling{
	// 	links = visit(links, c)
	// }

	return links
}