package main

func main() {

}

func permute(nums []int) [][]int {
	res := [][]int{}
	var dfs func(permutation []int, used []bool)

	dfs = func(permutation []int, used []bool) {
		if len(permutation) == len(nums) {
			c := make([]int, len(nums))
			copy(c, permutation)
			res = append(res, c)
			return
		}

		for i, num := range nums {
			if !used[i] {
				used[i] = true
				permutation = append(permutation, num)
				dfs(permutation, used)

				used[i] = false
				permutation = permutation[:len(permutation)-1]
			}
		}
	}

	dfs([]int{}, make([]bool, len(nums)))

	return res
}
