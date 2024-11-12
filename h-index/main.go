package main

import "sort"

func main() {

}

func hIndex(citations []int) int {
	sort.Ints(citations)
	h := 0
	length := len(citations)

	for i := range citations {
		if citations[i] >= length-i {
			h = length - i
			break
		}
	}

	return h
}
