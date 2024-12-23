package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"net/http"
	"golang.org/x/net/html"
)

type WordsImages struct{
	images map[string]int
	words []string

}
func main() {
	for _, url := range os.Args[1:] {
		word, image, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "finlinks2: %v\n", err)
			continue
		}
		fmt.Printf("Words: %d\nImages: %d\n", word, image)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	// Делаем запрос по переданному URL 
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	// Помещаем данные из тела запроса в переменную doc
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("failed HTML parsing: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)

	return
}

func countWordsAndImages(doc *html.Node) (words, images int) {
	// Объявляем мапу чтобы не вызвать панику
	res := WordsImages{
		images: make(map[string]int),
	}
	a := findKeys(res, doc)
	words = 0
	// Проходим массив a.words и в каждом элементе массива считает кол-во слов
	for _, word := range a.words{
		scanner := bufio.NewScanner(strings.NewReader(word))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan(){
			words++
		}
	}
	
	images = a.images["img"]
	
	return words, images
}

func findKeys(val WordsImages,doc *html.Node) WordsImages {
	 // Выполняем поиск тега <img> и подсчитываем кол-во с помощью карты
	if doc.Type == html.ElementNode {
		if doc.Data == "img"{
			for _, img := range doc.Attr{
				if img.Key == "src"{
					val.images[doc.Data]++
				}
			}
		}
	}
	// Выполняем поиск текста в запросе и отбрасываем теги текста, которые не видны на странице
	if doc.Type == html.TextNode && doc.Parent.Data != "script" && doc.Parent.Data != "style" {
        text := strings.TrimSpace(doc.Data)
		val.words = append(val.words, text)
        // fmt.Println(text)
    }
	if doc.FirstChild != nil {
		val = findKeys(val, doc.FirstChild)
	}
	if doc.NextSibling != nil {
		val = findKeys(val, doc.NextSibling)
	}
	return val
}

