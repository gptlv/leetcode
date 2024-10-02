package main

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	ptr1, ptr2, ptr := m-1, n-1, m+n-1

	for ptr1 >= 0 && ptr2 >= 0 && ptr >= 0 {
		if nums1[ptr1] >= nums2[ptr2] {
			nums1[ptr] = nums1[ptr1]
			ptr1--
		} else {
			nums1[ptr] = nums2[ptr2]
			ptr2--
		}

		ptr--
	}

	for ptr2 >= 0 {
		nums1[ptr] = nums2[ptr2]
		ptr--
		ptr2--
	}
}
