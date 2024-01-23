package main

func main() {

}

func canJump(nums []int) bool {
	idx := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		delta := idx - i

		if delta <= nums[i] {
			idx = i
		}

		if idx == 0 {
			return delta <= nums[i]
		}
	}

	return false
}

//https://leetcode.com/problems/jump-game/solutions/1285632/golang-o-n-o-1-solution-with-images/
