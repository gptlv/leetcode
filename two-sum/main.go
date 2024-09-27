package main

func main() {

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := []int{}

	for i := 0; i < len(nums); i++ {
		delta := target - nums[i]
		if _, exists := m[delta]; !exists {
			m[delta] = i
		}
	}

	for j := 1; j < len(nums); j++ {
		if i, exists := m[nums[j]]; exists && i != j {
			res = []int{i, j}
			break
		}
	}

	return res
}
