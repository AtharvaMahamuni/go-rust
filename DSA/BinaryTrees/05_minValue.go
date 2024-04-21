package main

func minValueDFS(root *NodeInt) int {
	if root == nil {
		return -1
	}

	stack := make([]*NodeInt, 1)
	stack[0] = root
	min := stack[0].Val

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Val < min {
			min = current.Val
		}

		if current.Left != nil {
			stack = append(stack, current.Left)
		}

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}

	return min
}

func helper(root *NodeInt, min int) int {

	if root == nil {
		return min
	}

	if root.Val < min {
		min = root.Val
	}

	leftMinValue := helper(root.Left, min)
	rightMinValue := helper(root.Right, min)

	if leftMinValue < rightMinValue && leftMinValue < min {
		min = leftMinValue
	}

	if rightMinValue < leftMinValue && rightMinValue < min {
		min = rightMinValue
	}

	return min
}

func minValueDFSRecursive(root *NodeInt) int {
	if root == nil {
		return -1
	}

	min := root.Val
	return helper(root, min)
}

func minValueBFS(root *NodeInt) int {
	if root == nil {
		return -1
	}

	queue := make([]*NodeInt, 1)
	queue[0] = root
	min := queue[0].Val

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if current.Val < min {
			min = current.Val
		}

		if current.Left != nil {
			queue = append([]*NodeInt{current.Left}, queue...)
		}

		if current.Right != nil {
			queue = append([]*NodeInt{current.Right}, queue...)
		}

	}

	return min
}
