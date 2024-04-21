package main

import "fmt"

func main() {

	// root := makeTree()
	root := makeIntTree()

	// fmt.Println("depthFirstIterative", depthFirstTraversal(root))
	// fmt.Println("depthFirstRecursive", depthFirstTraversalRecursive(root))
	// fmt.Println("breadthFirstItervative", breadthFirstTraversal(root))
	// fmt.Println("Tree Includes DFS", treeIncludesDFS(root, "y"))
	// fmt.Println("Tree Includes DFS recursive", treeIncludesDFSRecursive(root, "e"))
	// fmt.Println("Tree Includes DFS recursive", treeIncludesBFS(root, "z"))
	// fmt.Println("Tree sumDFS", treeSumDFS(root))
	// fmt.Println("Tree sumDFS Recursive", treeSumDFSRecursive(root, 0))
	// fmt.Println("Tree sumBFS", treeSumBFS(root))

	// fmt.Println("Min value DFS", minValueDFS(root))
	// fmt.Println("Min value DFS", minValueDFSRecursive(root))
	// fmt.Println("Min Value BFS", minValueBFS(root))

	fmt.Println("Max path sum DFS recursive", maxPathSumDFSRecursive(root))
	fmt.Println("Max path sum DFS", maxPathSumDFS(root))
}
