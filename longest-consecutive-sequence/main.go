package main

func main() {

}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]struct{})
	res, seq := 1, 1

	for _, num := range nums {
		m[num] = struct{}{}
	}

	startingPoints := []int{}

	for num := range m {
		var prev, next bool
		_, prev = m[num-1]
		_, next = m[num+1]

		if !prev && next {
			startingPoints = append(startingPoints, num)
		}
	}

	for _, p := range startingPoints {
		seq = countSequence(p, m)
		res = max(res, seq)
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func countSequence(startingPoint int, m map[int]struct{}) int {
	seq := 1

	for {
		if _, exists := m[startingPoint+1]; !exists {
			break
		}

		seq++
		startingPoint++
	}

	return seq
}
