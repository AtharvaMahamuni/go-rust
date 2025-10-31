package main

func treeSumDFS(root *NodeInt) int {
	if root == nil {
		return 0
	}

	sum := 0
	stack := make([]*NodeInt, 1)
	stack[0] = root

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += current.Val

		if current.Left != nil {
			stack = append(stack, current.Left)
		}

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}

	return sum
}

func treeSumDFSRecursive(root *NodeInt, sum int) int {
	if root == nil {
		return 0
	}

	leftValues := treeSumDFSRecursive(root.Left, sum)
	rightValues := treeSumDFSRecursive(root.Right, sum)

	return root.Val + leftValues + rightValues
}

func treeSumBFS(root *NodeInt) int {
	if root == nil {
		return 0
	}

	queue := make([]*NodeInt, 1)
	queue[0] = root
	sum := 0

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		sum += current.Val

		if current.Left != nil {
			queue = append([]*NodeInt{current.Left}, queue...)
		}

		if current.Right != nil {
			queue = append([]*NodeInt{current.Right}, queue...)
		}
	}

	return sum
}
