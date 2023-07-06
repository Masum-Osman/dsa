package array

import (
	"fmt"
	"sort"
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

/*
LC 242:
We can also use this snippet to count the frequency of a char in a desired string:
*/
func ValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	slice1, slice2 := [26]int{}, [26]int{}

	for i := 0; i < len(s); i++ {
		slice1[s[i]-'a']++
		slice2[t[i]-'a']++
	}

	return slice1 == slice2
}

/*
LC 1:
*/
func TwoSumWrong(nums []int, target int) []int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		sum := nums[start] + nums[end]
		if sum == target {
			return []int{start, end}
		} else {
			start++
			end--
		}
	}

	return []int{}
}

func TwoSum(nums []int, target int) []int {
	lookup := make(map[int]int)

	for i, v := range nums {
		j, ok := lookup[target-v]
		fmt.Println(j, ok)
		if ok {
			return []int{j, i}
		}

		lookup[v] = i
		fmt.Println("lookup map: ", lookup)
	}

	return []int{}
}

func SortString(s string) string {
	sChars := []rune(s)
	sort.Slice(sChars, func(i, j int) bool {
		return sChars[i] < sChars[j]
	})

	return string(sChars)
}

func GroupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	var ans [][]string

	for _, s := range strs {
		sorted := SortString(s)
		mp[sorted] = append(mp[sorted], s)
	}

	for _, v := range mp {
		ans = append(ans, v)
	}

	return ans
}

/*
vector<vector<string>> groupAnagrams(vector<string>& strs) {
    map<string, vector<string>> mp;
    vector<vector<string>> ans;

    // Iterate over the input vector of strings
    for(int i = 0; i < strs.size(); i++) {
        string s = strs[i];

        // Sort the current string in alphabetical order
        sort(strs[i].begin(), strs[i].end());

        // Add the original string to the corresponding anagram vector
        mp[strs[i]].push_back(s);
    }

    // Iterate over the map and add the anagram vectors to the result vector
    for(auto & it: mp) {
        ans.push_back(it.second);
    }

    return ans;
}
*/
