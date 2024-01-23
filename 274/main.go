package main

import "sort"

func main() {

}

func hIndex(citations []int) int {
	papers := len(citations)

	if papers == 0 {
		return 0
	}

	sort.Ints(citations)

	for i, citation := range citations {
		if citation >= papers-i {
			return papers - i
		}
	}

	return 0

}
