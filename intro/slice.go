package main

func SliceCopy(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	return res
}

func Clean(nums []int, x int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			for j := i; j < len(nums)-1; j++ {
				nums[j] = nums[j+1]
			}
			nums = nums[:len(nums)-1]
			i--
		}
	}
	return nums
}

func Join(nums1, nums2 []int) []int {
	res := make([]int, len(nums1)+len(nums2))
	copy(res, nums1)
	for i := 0; i < len(nums2); i++ {
		res[i+len(nums1)] = nums2[i]
	}
	return res
}

func Mix(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; 2*i < len(nums); i += 1 {
		res[2*i], res[2*i+1] = nums[i], nums[i+len(nums)/2]
	}
	return res
}
