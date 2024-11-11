package app

import (
	"unicode"
	"unicode/utf8"
)

func AppendInt(x []int, y int) []int {
	var z []int  // объявляем новый срез
	zlen := len(x) + 1 // добавляем к длине среза место под одно значение
	if zlen <= cap(x) { // проверям чтобы емкость среза была больше длины с доп. символом
		z = x[:zlen] // закидываем в z срез x от начала до новой длины
	} else {  // если длина с доп. символом больше емкости 
		zcap := zlen // новая емкость равна длине массива
		if zcap < len(x)*2 { //  если емкость меньше двойной длины
			zcap = len(x) * 2 // емкость равна двойной длине
		}
		z = make([]int, zlen, zcap)  // добавляем в новый срез
		copy(z, x) // копируем в новый  срез старый
	}
	z[len(x)] = y // в последний индекс закидываем y 
	return z
}

func Reverse(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func Rotate(s []int, k int) []int {
	n := len(s)
	temp := make([]int, n) // создаем новый срез
	for i := 0; i < n; i++ { 
		temp[i] = s[(k+i)%n] // i элемент среза = элементу из среза s по индексу (k+i)%n сдвиг на k индексов вправо, (k - i + n)%n k индексов влево 
	}
	return temp
}

func DeleteAdjValue(str []string) []string {
	index := 1   // индекс подсчета неравных элементов
	for i := 1; i < len(str)-1; i++ {  
		if str[i] != str[i-1] { // если элемент не равен предыдущему
			str[index] = str[i] // закидываем в массив по индексу, нужный элемент
			index++ // плюсуем индекс неравных элементов
		}
	}
	return str[:index] // обрезаем массив по длинне неравных элементув
}

func RemoveSpace(s []byte) []byte { 
	flag := false // ставим флаг изначально false чтобы зайти на иттерации с пробельным символом в условие
	pos := 0
	i := 0
	for i < len(s) {
		r, size := utf8.DecodeRune(s) // декодим срез байтов в руну
		if !unicode.IsSpace(r) { // если символ не пробел, то закинь этот символ обратно в срез 
			s[pos] = s[i] // 
			pos += size // плюсуем размер символа в байтах к индексу
			flag = false // ставим флаг false
		} else if !flag {  // если флаг false 
			copy(s[pos:], s[i:i+size]) // копируем в срез от pos до конца, срез от индекса до след символа
			pos += size // двигаем индекс на кол-во байт символа
			flag = true // 
		}
		i += size // двигаем индекс на количество байт
	}
	return s[:pos]
}

func NewReverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i,j = i+1, j-1{
		s[i], s[j] = s[j], s[i]
	}
	return s
}