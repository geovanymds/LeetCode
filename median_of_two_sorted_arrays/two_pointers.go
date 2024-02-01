func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums3 []int
	m, n := len(nums1), len(nums2)
	i, j := 0, 0

	for i < m && j < n && i+j <= (m+n)/2 {
		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i++
		} else {
			nums3 = append(nums3, nums2[j])
			j++
		}
	}

	for i < m && (i+j) <= (m+n)/2 {
		nums3 = append(nums3, nums1[i])
		i++
	}

	for j < n && (i+j) <= (m+n)/2 {
		nums3 = append(nums3, nums2[j])
		j++
	}

	if (m+n)%2 == 0 {
		return (float64(nums3[len(nums3)-1]) + float64(float64(nums3[len(nums3)-2]))) / float64(2)
	} else {
		return float64(nums3[len(nums3)-1])
	}

}