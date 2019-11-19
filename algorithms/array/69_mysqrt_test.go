package array

import (
	"fmt"
	"testing"
)

func TestMysqrt(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 8, 9, 16, 100}
	//nums := []int{9}

	for _, x := range nums {
		fmt.Println(mySqrt(x))
	}
}
