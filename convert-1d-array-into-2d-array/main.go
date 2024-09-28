package main

func main() {

}

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	res := make([][]int, m)

	for ptr := 0; ptr < len(original); ptr++ {
		res[ptr/n] = append(res[ptr/n], original[ptr])
	}

	return res
}
