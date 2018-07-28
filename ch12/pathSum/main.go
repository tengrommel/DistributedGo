package main

/*
给定一个二叉树与整数sum，找出所有从根节点到叶节点的路径，这些路径上的节点值累加和为sum
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	return helper(root, sum, []int{})
}

func helper(root *TreeNode, sum int, solution []int) (rets [][]int) {
	if root == nil{
		return
	}
	if root.Left == nil && root.Right == nil{
		if root.Val != sum{
			return
		}
		tmp := make([]int, len(solution)+1)
		copy(tmp, solution)
		tmp[len(solution)] = sum
		rets = append(rets, tmp)
	}
	solution = append(solution, root.Val)
	leftRets := helper(root.Left, sum-root.Val, solution)
	rightRets := helper(root.Right, sum-root.Val, solution)
	if len(leftRets) != 0 {
		rets = append(rets, leftRets...)
	}
	if len(rightRets) != 0 {
		rets = append(rets, rightRets...)
	}
	return rets
}

func main() {
	
}
