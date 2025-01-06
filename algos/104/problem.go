package algos

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	depth := 0
	var dfs func(t *TreeNode, d int)
	dfs = func(t *TreeNode, d int) {
		if t == nil {
			depth = max(depth, d)
			return
		}
		dfs(t.Left, d+1)
		dfs(t.Right, d+1)
	}
	dfs(root, 0)
	return depth
}

func AltMaxDepth(root *TreeNode) int {
	return height(root)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	hl := height(root.Left)
	hr := height(root.Right)
	if hl < hr {
		return hr + 1
	}
	return hl + 1
}
