package tree

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

func lowestCommonAncestor_235(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
