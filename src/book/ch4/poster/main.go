package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"

	"bookProjects/ch4/poster/posterPack"
)

func main() {

	var titleFilm string

	reader := bufio.NewReader(os.Stdin)
	titleFilm, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	titleFilm = strings.TrimSpace(titleFilm)
	formatTitle := strings.ReplaceAll(titleFilm, " ", "+")

	req, err := posterPack.PosterRequest(formatTitle)
	if err != nil {
		return
	}

	if len(req.Search) > 1 {
		fmt.Println("ПО ВАШЕМУ ЗАПРОСУ НАЙДЕНО ФИЛЬМОВ:")
		for i, title := range req.Search {
			fmt.Println(i, i+1, title.Title)
		}
		fmt.Println("ВВЕДИТЕ НОМЕР ФИЛЬМА:")
		var filmId int
		fmt.Scan(&filmId)
		if filmId > 0 && filmId <= len(req.Search) {
			fmt.Printf("Название:\t%s\nURL Афиши:\t%s\n", req.Search[filmId-1].Title, req.Search[filmId-1].Poster)
		} else {
			fmt.Println("ВЫ ВВЕЛИ НЕВЕРНЫЙ НОМЕР")
			return
		}
	} else {
		fmt.Printf("Название:\t%s\nURL Афиши:\t%s\n", req.Search[0].Title, req.Search[0].Poster)

	}
}
