package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	sum := numbers[i] + numbers[j]

	for sum != target {
		if sum > target {
			j--
		}

		if sum < target {
			i++
		}

		sum = numbers[i] + numbers[j]
	}

	return []int{i + 1, j + 1}
}
