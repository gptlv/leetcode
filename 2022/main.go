package main

func main() {

}

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	res := make([][]int, m)

	for i := 0; i <= len(original)-n; i += n {
		res[i/n] = original[i : i+n]
	}

	return res
}
