package main

func main() {

}

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}

	left := 0
	right := k - 1
	max := avg(nums[0:k])
	average := avg(nums[0:k])

	left++
	right++

	for right < len(nums) {
		average = average + float64(nums[right]-nums[left-1])/float64(k)
		if average > max {
			max = average
		}

		right++
		left++
	}

	return max
}

func avg(nums []int) float64 {
	var avg float64

	for _, num := range nums {
		avg += float64(num)
	}

	return avg / float64(len(nums))
}
