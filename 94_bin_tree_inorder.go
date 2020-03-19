//
// Snippets only
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	r := new([]int)
	tr(root, r)
	return *r
}

func tr(t *TreeNode, o *[]int) {
	if t == nil {
		return
	}

	tr(t.Left, o)
	*o = append(*o, t.Val)
	tr(t.Right, o)
}
