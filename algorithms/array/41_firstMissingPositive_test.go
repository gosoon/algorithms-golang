package array

import (
	"fmt"
	"testing"
)

// TestFirstMissingPositive  xxx
func TestFirstMissingPositive(t *testing.T) {
	nums := [][]int{
		{1, 2, 0},
		{3, 4, -1, 1},
		{7, 8, 9, 11, 12},
	}

	for _, num := range nums {
		fmt.Println(firstMissingPositive(num))
	}
}
