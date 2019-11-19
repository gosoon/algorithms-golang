package array

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left := 0
	right := x
	mid := (left + right) / 2

	for left < right {
		//fmt.Println(left, right)
		t := mid * mid
		if t == x {
			return mid
		}
		if left+1 == right {
			return left
		}

		if t > x {
			right = mid
		} else {
			left = mid

		}
		mid = (right + left) / 2
	}
	return 0
}
