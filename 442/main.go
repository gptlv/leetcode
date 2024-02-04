package main

func main() {

}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func findDuplicates(nums []int) []int {
	res := []int{}

	for _, num := range nums {
		if nums[abs(num)-1] < 0 {
			res = append(res, abs(num))
		} else {
			nums[abs(num)-1] = -1 * nums[abs(num)-1]
		}
	}

	return res

	// initial solution O(n) space
	//
	// m := make(map[int]struct{})
	// idx := 0

	// for _, num := range(nums) {
	//     if _, ok := m[num]; ok {
	//         nums[idx] = num
	//         idx++
	//     }
	//     m[num] = struct{}{}
	// }

	// return nums[:idx]
}
