/* https://leetcode.com/problems/longest-common-prefix/

Write a function to find the longest common prefix string amongst an array of
strings.

If there is no common prefix, return an empty string "".
*/
package main

func longestCommonPrefix(strings []string) string {
	switch len(strings) {
	case 0:
		return ""
	case 1:
		return strings[0]
	}
	result := strings[0]
	for index := 1; index < len(strings); index++ {
		current := strings[index]
		previous := strings[index-1]
		current_max := len(current)
		result_max := len(result)
		var slice_max int
		if result_max > current_max {
			slice_max = current_max
			result = result[0:slice_max]
		} else {
			slice_max = result_max
		}
		for ; slice_max >= 0; slice_max-- {
			current_slice := current[0:slice_max]
			if current_slice == previous[0:slice_max] {
				result = current_slice
				break
			}
			if slice_max == 0 {
				return ""
			}
		}
	}
	return result
}

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
	longestCommonPrefix([]string{"dog", "racecar", "car"})
	longestCommonPrefix([]string{"a"})
	longestCommonPrefix([]string{"cir", "car"})
}
