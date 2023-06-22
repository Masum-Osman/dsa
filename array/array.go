package array

import (
	"fmt"
	"reflect"
)

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

// LC 217
func ContainsDuplicateUsingMap(nums []int) bool {
	nums_map := map[int]int{}

	for _, n := range nums {
		if _, ok := nums_map[n]; !ok {
			nums_map[n] = 1
		} else {
			return true
		}
	}
	return false
}

func ReverseString(s string) string {
	output := ""

	for i := 1; i <= len(s); i++ {
		output += string(s[len(s)-i])
	}
	return output
}

func ValidAnagramBruteForce(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	map1 := map[int]byte{}
	map2 := map[int]byte{}

	for i := 0; i < len(s); i++ {
		map1[i] = s[i]
		map2[i] = t[i]
	}

	fmt.Println(map1, map2)

	return reflect.DeepEqual(map1, map2)
}
