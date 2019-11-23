package array

// https://leetcode-cn.com/problems/trapping-rain-water/

func trap(height []int) int {
	ant := 0
	l := len(height)
	for t, val := range height {
		// right max
		rightMax := val
		for j := t; j < l; j++ {
			if height[j] > rightMax {
				rightMax = height[j]
			}
		}

		// left max
		leftMax := val
		for i := t; i >= 0; i-- {
			if height[i] > leftMax {
				leftMax = height[i]
			}
		}

		// min(rightMax,leftMax)
		if rightMax > leftMax {
			ant += leftMax - height[t]
		} else {
			ant += rightMax - height[t]
		}
	}
	return ant
}
