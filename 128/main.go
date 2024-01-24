package main

func main() {

}

func longestConsecutive(nums []int) int {

	m := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		m[num] = struct{}{}
	}

	longest := 0

	for n := range m {
		if _, ok := m[n-1]; !ok {
			length := 1
			for _, ok := m[n+length]; ok; _, ok = m[n+length] {
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest

}
