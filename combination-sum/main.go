package main

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func(int, []int, int)

	dfs = func(i int, combination []int, sum int) {
		if sum == target {
			c := make([]int, len(combination))
			copy(c, combination)
			res = append(res, c)
			return
		}

		if sum > target || i >= len(candidates) {
			return
		}

		combination = append(combination, candidates[i])
		dfs(i, combination, sum+candidates[i])

		combination = combination[:len(combination)-1]

		dfs(i+1, combination, sum)
	}

	dfs(0, []int{}, 0)

	return res
}
