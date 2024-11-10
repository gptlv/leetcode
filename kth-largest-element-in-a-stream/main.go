package main

func main() {

}

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, num := range nums {
		kl.Add(num)
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) < this.k {
		this.nums = append(this.nums, val)
		this.upHeap(len(this.nums) - 1)
	} else if val > this.nums[0] {
		this.nums[0] = val
		this.downHeap(0)
	}

	return this.nums[0]

}

func (this *KthLargest) upHeap(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if this.nums[index] >= this.nums[parent] {
			break
		}

		this.swap(index, parent)
		index = parent
	}
}

func (this *KthLargest) downHeap(index int) {
	n := len(this.nums)
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < n && this.nums[left] < this.nums[smallest] {
			smallest = left
		}

		if right < n && this.nums[right] < this.nums[smallest] {
			smallest = right
		}

		if smallest == index {
			break
		}

		this.swap(index, smallest)
		index = smallest
	}
}

func (this *KthLargest) swap(i, j int) {
	this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
