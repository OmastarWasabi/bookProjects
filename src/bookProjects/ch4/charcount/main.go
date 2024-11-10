package main 

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int) // создаем мапу с ключом рун и значением инт

	letterCount := 0
	digitCount := 0
	otherCount := 0
	var utflen [utf8.UTFMax + 1]int  // создаем массив длинной utf8.UTFMax = 4 + 1, сколько символов имеет байт от 1 до 4
	invalid := 0 // счетчик некоректных символов utf-8
	
	in := bufio.NewReader(os.Stdin) // считываем значения из ввода в буфер
	for {
		r, n, err := in.ReadRune() // читает руну и возвращает руну, байты и ошибки 
		if err == io.EOF { // если ошибка равна концу файла или ввода выйди
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1) //если ошибка есть, то выведи в поток ошибок зничение ошибки и выйди из программы с кодом 1
		}
		if r == unicode.ReplacementChar && n == 1 { // если руна равна спец символу и кол-во байт 1 прибавь к счетчику 1
			invalid++
			continue
		}
		if unicode.IsLetter(r){
			letterCount++
		} else if unicode.IsDigit(r){
			digitCount++
		} else {
			otherCount++
		}
		counts[r]++
		utflen[n]++ // +1 к значению по индексу кол-ва байт, которое занимает символ
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)  // пробежимся по отображению рун и позначениям(сколько раз была найдена руна)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen{
		if  i > 0{
		fmt.Printf("%d\t%d\n", i, n) //выводим индекс(кол-во байт) и значение по индексу
		} 
	}
	fmt.Printf("Букв: %d Цифр:%d Прочие:%d", letterCount, digitCount, otherCount)
	if invalid > 0 {
		fmt.Printf("\n%d неверных символов UTF-8\n", invalid) // если есть неверные символы выводим количество
	}

}


