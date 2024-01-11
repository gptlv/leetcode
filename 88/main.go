package main

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)

}
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		copy(nums1, nums2)

		return
	}

	pt1 := m - 1
	pt2 := n - 1
	pt3 := n + m - 1

	for pt1 >= 0 && pt2 >= 0 {
		if nums1[pt1] > nums2[pt2] {
			nums1[pt3] = nums1[pt1]
			pt1--
		} else {
			nums1[pt3] = nums2[pt2]
			pt2--
		}
		pt3--
	}

	if pt2 >= 0 {
		copy(nums1[:pt3+1], nums2[:pt2+1])
	}

}
