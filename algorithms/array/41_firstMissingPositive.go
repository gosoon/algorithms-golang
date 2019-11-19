package array

// https://leetcode-cn.com/problems/first-missing-positive/
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	minInt := 1
	m := make(map[int]int)
	for _, n := range nums {
		if _, found := m[n]; !found {
			m[n] = n
		}

		if _, found := m[minInt]; found {
			minInt++
		}

	}

	for i := 0; i < len(nums); i++ {
		if _, found := m[minInt]; found {
			minInt++
		}
	}
	return minInt
}
