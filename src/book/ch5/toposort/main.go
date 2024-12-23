package main

import (
	"sort"
	"fmt"
)

func topoSort(m map[string][]string)[]string{
	var order []string
		seen := make(map[string]bool)
		var visitAll func(items []string)
		visitAll = func(items []string) {
			for _, item := range items{
				if !seen[item]{
					seen[item]=true
					visitAll(m[item])
					order = append(order, item)
				}
			}
		}
		var keys []string
		for key := range m{
			keys = append(keys, key)
		}
		sort.Strings(keys)
		visitAll(keys)
		return order
}


func main(){
	var prereqs = map[string][]string{
		"algorithms":{"data structures"},
		"calculus": {
			"data structures",
			"formal languages",
			"computer organization",
			},
		"data structures": {"discrete math"},
		"databases": {"data structures"},
		"discrete math": {"intro to programming"},
		"formal languages": {"discrete math"},
		"networks": {"operating system"},
		"operating system": {"data structures",
							"computer organization"},
		"programming languages": {"data structures",
								  "computer organization"},					
	}
	for i, course := range topoSort(prereqs){
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

