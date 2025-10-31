package main

func depthFirstTraversal(root *Node) []string {
	result := make([]string, 0)

	if root == nil {
		// fmt.Println(result)
		return result
	}

	stack := make([]*Node, 1)
	stack[0] = root

	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		if current.Right != nil {
			stack = append(stack, current.Right)
		}

		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	// fmt.Println(result)
	return result
}

func depthFirstTraversalRecursive(root *Node) []string {

	if root == nil {
		return make([]string, 0)
	}

	leftValues := depthFirstTraversalRecursive(root.Left)
	rightValues := depthFirstTraversalRecursive(root.Right)

	value := make([]string, 1)
	value[0] = root.Val
	return append(value, append(leftValues, rightValues...)...)
}
