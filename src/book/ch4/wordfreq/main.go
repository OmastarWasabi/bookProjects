package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2{
		fmt.Printf("Чтобы запустить программу введите: youprog.go <имя файла>")
		os.Exit(1)
	}
	input := os.Args[1]
	file, err := os.Open(input)
		if err != nil{
			fmt.Fprintf(os.Stderr, "Ошибка открытия файла: %v", err)
			os.Exit(1)
		}
	defer file.Close()
	words := make(map[string]int)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words[scanner.Text()]++
	}
	fmt.Printf("Word\tcount\n")
	for word, count := range words{
		fmt.Printf("%s\t%d\n", word, count)
	}
}