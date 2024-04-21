package main

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
