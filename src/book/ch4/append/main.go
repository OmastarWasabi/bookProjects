package main

import (
	"fmt"

	"bookProjects/ch4/append/app"
)

func main() {
    b := []byte("12345")
	fmt.Println(string(app.NewReverse(b)))
}



