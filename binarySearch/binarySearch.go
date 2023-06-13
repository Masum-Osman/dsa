package binarySearch

func Search(nums []int, target int) int {
	left := 0
	right := len(nums)

	for i := left; i < right; i++ {
		if target == nums[i] {
			return i
		} else if target > nums[i] {
			left = (left + right) / 2
		} else {
			right = (left + right) / 2
		}
	}

	return -1
}

func SearchInsert(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		// fmt.Println("LEFT: ", left, "RIGHT: ", right, "\t", "MID: ", mid, "\t",
		// 	"NUMS[MIS]: ", nums[mid])

		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func SearchInsertUpdated(nums []int, target int) int {
	index := 0
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
			// index = r
		} else {
			l = mid + 1
			index = l
		}
	}

	return index
}
