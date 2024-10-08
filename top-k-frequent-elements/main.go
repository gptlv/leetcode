package main

func main() {
	topKFrequent([]int{1, 2}, 2)
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	arr := make([][]int, len(nums)+1)
	res := []int{}

	for _, num := range nums {
		m[num]++
	}

	for key, value := range m {
		arr[value] = append(arr[value], key)
	}

	found := 0
	ptr := len(arr) - 1

	for len(res) < k && ptr >= 0 {
		if len(arr[ptr]) != 0 {
			found++
			res = append(res, arr[ptr]...)
		}
		ptr--
	}

	return res
}
