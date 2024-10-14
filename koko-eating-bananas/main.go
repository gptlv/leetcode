package main

func main() {

}

func minEatingSpeed(piles []int, h int) int {
	max := getMax(piles)

	left := 1
	right := max

	res := max

	for left <= right {
		mid := (left + right) / 2

		hours := getHours(piles, mid)

		if hours <= h {
			res = min(res, mid)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res

}

func getHours(piles []int, speed int) int {
	t := 0

	for _, pile := range piles {
		t += (pile + speed - 1) / speed
	}

	return t
}

func getMax(piles []int) int {
	max := piles[0]

	for i := 1; i < len(piles); i++ {
		if piles[i] > max {
			max = piles[i]
		}
	}

	return max
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
