package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Join("...", "a", "b", "c"))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "..."))

}

func min(vals ...int) int {
	min := vals[0]
	if len(vals) > 0 {
		for _, val := range vals {
			if val < min {
				min = val
			}
		}
	} else {
		return 0
	}
	return min
}

func max(vals ...int) int {
	max := vals[0]
	if len(vals) > 0 {
		for _, val := range vals {
			if val > max {
				max = val
			}
		}
	} else {
		return 0
	}
	return max
}

func Join(sep string, strs ...string) string{
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}

	result := strs[0]

	for i := 1; i < len(strs); i++{
		result += sep + strs[i]
	}
	return result
}
