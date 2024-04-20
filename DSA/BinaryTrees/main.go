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
	fmt.Println("Tree sum", treeSumDFS(root))
}
