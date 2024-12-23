package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"book/packages"
)

func main() {
	year, month, _ := time.Now().Date()
	YearNow := formatYear(year, month)

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d тем:\n", result.TotalCount)
	for _, item := range result.Items {
		year, month, _ := item.CreatedAt.Date()
		YearItem := formatYear(year, month)
		resString := ""
		if YearNow-YearItem < 1 {
			resString = "Менее месяца назад"
		} else {
			if YearNow-YearItem < 12 {
				resString = "Менее года назад"
			} else {
				resString = "Более года назад"
			}
		}

		fmt.Printf("#%-5d %9.9s %.55s давность: %s; %d %d\n",
			item.Number, item.User.Login, item.Title, resString, year, int(month))
	}
}

func formatYear(year int, month time.Month) int {
	YearNow := year*12 + int(month)
	return YearNow
}
