package array

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	//s := []string{"flower", "flow", "flight"}
	s := []string{"aa", "a"}
	fmt.Println(longestCommonPrefix(s))
}
