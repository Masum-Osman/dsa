package array

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	/*
		for i := 0; i < 5; i++ {
			sum += numbers[i]
		}
	*/

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func ContainsDuplicateTwoPointer(nums []int) bool {

	start := 0
	end := len(nums) - 1

	for start <= end {
		// fmt.Println(nums[start], nums[end])
		fmt.Println(start, end)
		if nums[start] == nums[end] && len(nums) != 1 && start != end {
			return true
		} else {
			start++
			end--
		}
	}

	return false
}

func ContainsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func ContainsDuplicateUsingMap(nums []int) bool {
	nums_map := make(map[int]int)
	nums_map[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		val, exist := nums_map[nums[i]]
		fmt.Println(val, exist)
	}

	return false
}
