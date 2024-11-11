package main

import (
	"bookProjects/ch4/append/app"
	"fmt"
	// "unicode"
	// "unicode/utf8"
	// "bookProjects/ch4/append/app"
)

func main() {
    b := []byte("12345")
	fmt.Println(string(app.NewReverse(b)))
    // for len(b) > 0 {

    //     r, size := utf8.DecodeRune(b)
		
    //     fmt.Printf("%c  size:%d len: %d ", r, size, len(b)) // Выводит каждый символ (руну)
    //     b = b[size:] 
	// 	fmt.Printf("%s\n", b)        // Продвигаем срез байтов
    // }
}



