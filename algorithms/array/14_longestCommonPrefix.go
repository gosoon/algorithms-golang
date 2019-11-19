package array

// https://leetcode-cn.com/problems/longest-common-prefix/submissions/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var comm string
	for i := 0; i < len(strs[0]); i++ {
		f := strs[0][i]
		flag := true
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 || strs[j][i] != f {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		comm = strs[0][0 : i+1]
	}

	return comm
}
