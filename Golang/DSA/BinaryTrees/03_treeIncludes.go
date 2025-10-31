package main

func treeIncludesDFS(root *Node, target string) bool {

	if root == nil {
		return false
	}

	stack := make([]*Node, 1)
	stack[0] = root

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Val == target {
			return true
		}

		if current.Left != nil {
			stack = append(stack, current.Left)
		}

		if current.Right != nil {
			stack = append(stack, current.Right)
		}

	}

	return false
}

func treeIncludesDFSRecursive(root *Node, target string) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}

	leftValues := treeIncludesDFSRecursive(root.Left, target)
	rightValues := treeIncludesDFSRecursive(root.Right, target)

	return leftValues || rightValues
}

func treeIncludesBFS(root *Node, target string) bool {
	if root == nil {
		return false
	}

	queue := make([]*Node, 1)
	queue[0] = root

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if current.Val == target {
			return true
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return false
}
