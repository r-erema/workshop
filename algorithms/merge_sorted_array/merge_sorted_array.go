package mergesortedarray

// Merge merges two sorted arrays where nums1 has enough space for both.
// Time O(m+n), since we iterate once both arrays
//
// Space O(1), since we don't involve any extra space.
func Merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[n+m-1] = nums1[m-1]
			m--
		} else {
			nums1[n+m-1] = nums2[n-1]
			n--
		}
	}
}
