import "slices"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var merged_array []int

	merged_array = append(merged_array, nums1...)
	merged_array = append(merged_array, nums2...)

	slices.Sort(merged_array)

	median := float64(0.0)

	if len(merged_array)%2 != 0 {
		median = float64(merged_array[len(merged_array)/2])
	} else {
		median = (float64(merged_array[len(merged_array)/2]) + float64(merged_array[(len(merged_array)/2)-1])) / float64(2)
	}

	return median

}