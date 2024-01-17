package main

func main() {

	removeDuplicates([]int{1, 1, 1, 2, 2, 3})

}

func removeDuplicates(nums []int) int {
	m := make(map[int]int)

	idx := 0

	for i, num := range nums {
		m[num]++

		if m[num] == 3 && idx == 0 {
			idx = i
			continue
		}

		if m[num] > 2 {
			continue
		}

		nums[idx] = nums[i]
		idx++
	}

	return idx
}
