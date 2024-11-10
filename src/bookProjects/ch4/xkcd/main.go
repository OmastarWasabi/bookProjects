package main

import (
	"bookProjects/ch4/xkcd/xkcd"
	"fmt"
	"os"
)

func main() {
	number := os.Args[1]
	request, err := xkcd.ComicsRequest(number)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := xkcd.ComicsWriteFile(*request)
	res := result[request.Num]
	
    fmt.Printf("URL: %s\nDescription: %s", res.ImgURL, res.Description)
}
