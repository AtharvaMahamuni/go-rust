package main

func maxPathSumDFS(root *NodeInt) int {
	if root == nil {
		return 0
	}

	stack := make([]*NodeInt, 1)
	stack[0] = root
	maxSum := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Left != nil && current.Right != nil {
			if current.Left.Val > current.Right.Val {
				maxSum += current.Left.Val
			}

			if current.Right.Val > current.Left.Val {
				maxSum += current.Right.Val
			}
		}
	}

	return maxSum
}

func helperSum(root *NodeInt, sum int) int {
	if root == nil {
		return 0
	}

	sum += root.Val
	leftValueSums := helperSum(root.Left, sum)
	rightValueSums := helperSum(root.Right, sum)

	if leftValueSums > rightValueSums {
		return leftValueSums
	}

	if rightValueSums > leftValueSums {
		return rightValueSums
	}

	return sum
}

func maxPathSumDFSRecursive(root *NodeInt) int {
	if root == nil {
		return 0
	}

	return helperSum(root, 0)
}
