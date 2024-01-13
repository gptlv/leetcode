package main

import "fmt"

func main() {

	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))

}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{fmt.Sprintf("%v", nums[0])}
	}

	res := []string{}

	var p1 int
	var p2 int

	for p2 < len(nums)-1 {
		if nums[p2]+1 == nums[p2+1] {
			p2++
			continue
		}

		if p1 == p2 {
			res = append(res, fmt.Sprintf("%v", nums[p1]))
		} else {
			res = append(res, fmt.Sprintf("%v->%v", nums[p1], nums[p2]))
		}

		p2++
		p1 = p2
	}

	if p1 == p2 {
		res = append(res, fmt.Sprintf("%v", nums[p1]))
	} else {
		res = append(res, fmt.Sprintf("%v->%v", nums[p1], nums[p2]))
	}

	return res
}
