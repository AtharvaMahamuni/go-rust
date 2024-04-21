package main

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type NodeInt struct {
	Val   int
	Left  *NodeInt
	Right *NodeInt
}

func makeIntTree() *NodeInt {
	a := NodeInt{Val: 5}
	b := NodeInt{Val: 8}
	c := NodeInt{Val: 9}
	d := NodeInt{Val: 3}
	e := NodeInt{Val: 21}
	f := NodeInt{Val: 12}

	a.Left = &b
	a.Right = &c
	b.Left = &d
	b.Right = &e
	c.Right = &f

	return &a
}

func makeTree() *Node {

	a := Node{Val: "a"}
	b := Node{Val: "b"}
	c := Node{Val: "c"}
	d := Node{Val: "d"}
	e := Node{Val: "e"}
	f := Node{Val: "f"}

	a.Left = &b
	a.Right = &c
	b.Left = &d
	b.Right = &e
	c.Right = &f

	//  	   a
	//        / \
	//       b   c
	//      / \   \
	//     d   e   f

	return &a
	// return nil
}
