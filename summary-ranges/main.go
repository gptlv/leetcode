package main

import (
	"fmt"
)

func main() {

}

func summaryRanges(nums []int) []string {
    res := []string{}
    store := []int{}

    if len(nums) == 0 {
        return res
    }

	for i := range nums {
		if i == 0 || nums[i-1] == nums[i]-1 {
			store = append(store, nums[i])
			continue
		}

		res = append(res, getRange(store))
		store = nil
		store = append(store, nums[i])
	}

	return append(res, getRange(store))
}

func getRange(store []int) string {
	if len(store) == 0 {
		return ""
	}

	if len(store) == 1 {
		return fmt.Sprintf("%d", store[0])
	}

	return fmt.Sprintf("%d->%d", store[0], store[len(store)-1])
}
