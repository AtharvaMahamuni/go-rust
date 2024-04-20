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
