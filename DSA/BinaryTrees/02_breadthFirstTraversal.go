package main

func breadthFirstTraversal(root *Node) []string {
	result := make([]string, 0)

	if root == nil {
		return result
	}

	queue := make([]*Node, 1)
	queue[0] = root

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		result = append(result, current.Val)

		if current.Left != nil {
			queue = append([]*Node{current.Left}, queue...)
		}

		if current.Right != nil {
			queue = append([]*Node{current.Right}, queue...)
		}
	}

	return result
}

// no recursion for breadth first
// it is difficult to maintain queue, in depth first we 
// maintain the os queue
