package array

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	testCases := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{2, 0, 2},
	}

	for _, test := range testCases {
		fmt.Println(trap(test))
	}
}
