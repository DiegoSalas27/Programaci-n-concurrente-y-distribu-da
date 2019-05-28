package main

import (
	"fmt"
)

type Node struct {
	Left *Node
	Value int
	Right *Node
}

func inOrder(n *Node, ch chan int) {
	if n != nil {
		inOrder(n.Left, ch)
		ch <- n.Value
		inOrder(n.Right, ch)
	}
}

func add(n *Node, e int) *Node {
	if n == nil {
		return &Node{Value: e}
	} else if n.Value < e {
		n.Right = add(n.Right, e)
	} else {
		n.Left = add(n.Left, e)
	}
	return n
}

func check(r1, r2 *Node) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go inOrder(r1, ch1)
	go inOrder(r2, ch2)
	for i := 0; i < 5; i++  {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	var root *Node
	root = add(root, 5)
	root = add(root, 3)
	root = add(root, 4)
	root = add(root, 1)
	root = add(root, 2)

	fmt.Println(root.Value)

	var root2 *Node
	root2 = add(root2, 1)
	root2 = add(root2, 2)
	root2 = add(root2, 3)
	root2 = add(root2, 4)
	root2 = add(root2, 5)

	if check(root, root2) {
		fmt.Println(":)")
	} else {
		fmt.Println(":(")
	}
}