package main

import (
	"os"
	"crypto/sha512"
	"fmt"
	
	
)

func main() {
    if len(os.Args) < 2{
		fmt.Println("Введите: go run main.go <value> [S384||S512] или <value> <value>")
		return
	}

	value := os.Args[1]
	if len(os.Args) == 2 {
		digest := sha512.Sum512_256([]byte(value))
		fmt.Printf("SHA256: %x; \n", digest)
		return
	}
	if len(os.Args) == 3 {
		value2 := os.Args[2]
		digest1 := sha512.Sum512_256([]byte(value))
		digest2 := sha512.Sum512_256([]byte(value2))
		differentsBits:= countDiffBits(digest1, digest2)
		fmt.Printf("разных битов в SHA256: %x; \n", differentsBits)
		return
	}

	
	flag := os.Args[2]

	
	switch flag {
	case "S384":
		fmt.Printf("SHA384: %x; \n", sha512.Sum384([]byte(value)))
	case "S512":
		fmt.Printf("SHA512: %x; \n", sha512.Sum512([]byte(value)))
	default:
		fmt.Printf("%s", "Введи S348, либо S512, ЕБЛАН!")
	}	

}


// функция для подсчета различных битов
func countDiffBits(digest1, digest2 [32]byte) int{
	total:= 0 // счетчик битов
	for i:= 0; i < len(digest1); i++{ 
		total += countBits(digest1[i]^digest2[i]) // ^ - побитовое исключение 101 ^ 100 = 011
	}
	return total
}

// считает единичные биты
func countBits(b byte) int {  
	total := 0
		for b > 0{
			total += int(b & 1)   // int(11011111 & 1 = true и т.д)
			b >>= 1 // сдвигаем биты на 1 право
		}
	return total	
}

// функция из главы 2.6.2

// func bitsHash(hash []byte) int {
// 	total := 0
// 	for _, v := range hash {
// 		total += PopCount(uint64(v))
// 	}
// 	return total
// }

// func PopCount(x uint64) int {
// 	totalPc := 0
// 	bit := x
// 	for bit != 0 {
// 		bit = bit & (bit - 1) 
// 		totalPc++

// 	}
// 	return totalPc
// }
