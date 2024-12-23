package comm

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func Anagramm(str1, str2 string) bool {
	str1= strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if len(str1) != len(str2) {
		return false
	}

	runes1 := []rune(str1)
	runes2 := []rune(str2)

	sort.Slice(runes1, func (i, j int) bool {return runes1[i] < runes1[j]})
	sort.Slice(runes2, func (i, j int) bool {return runes2[i] < runes2[j]})


	return string(runes1) == string(runes2)
}

func Comma(s string) string {
	var sign string
	if s[0] == '+' || s[0] == '-'{
		sign = string(s[0])
		s = s[1:]
	}
	indexDot := DigitFraction(s)
	var dig,  frac string
	if indexDot > 0{
		dig = s[:indexDot]
		frac = s[indexDot:]
	} else {
		dig = s
	}
	n := len(dig)
	interval := n % 3
	var buf bytes.Buffer 
	for i, ch := range dig {
		if i > 0 && (i - interval) % 3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteString(string(ch))

	}
	
	return sign + buf.String() + frac
}

func DigitFraction(str string) int {
	flag := 0
	for i, ch := range str {
		if ch == '.' {
			flag = i
			break
		}
		
	}
	return flag
}

func Factorial(n int) int {

	if n == 1 {
		return 1
	}
	return Factorial(n-1) * (n)
}

func Fibonacci(n int) {
	num1, num2 := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(num1)
		res := num1 + num2
		num1, num2 = num2, res

	}

}
