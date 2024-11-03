package main

func main() {

}

func subsets(nums []int) [][]int {
	res := [][]int{{}}

	for _, num := range nums {
		length := res
		for i := range length {
			subset := append([]int{}, res[i]...)
			subset = append(subset, num)

			res = append(res, subset)
		}

	}

	return res
}
