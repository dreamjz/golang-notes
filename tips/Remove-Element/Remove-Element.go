package tips

// https://leetcode-cn.com/leetbook/read/array-and-string/cv3bv/

func removeElementWithExtraArray(nums []int, val int) int {
	var arr []int
	for _, v := range nums {
		if v != val {
			arr = append(arr, v)
		}
	}
	return len(arr)
}

func removeElementDoublePointer(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
