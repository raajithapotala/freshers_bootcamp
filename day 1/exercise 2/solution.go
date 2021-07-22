package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func treeCeator(st []string) *Node {
	if len(st) == 0 {
		return nil
	}
	if !checkoperator(st[0]) {
		node := Node{st[0], nil, nil}
		head := treeCeator(st[1:])
		if head == nil {
			return &node
		}
		head.left = &node
		return head
	} else {
		node := Node{st[0], nil, nil}
		right := treeCeator(st[1:])
		node.right = right
		return &node
	}

}
func checkoperator(s string) bool {
	if s == "+" || s == "-" {
		return true
	}
	return false
}

func preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.data + " ")
	preorder(node.left)
	preorder(node.right)

}

func postorder(node *Node) {
	if node == nil {
		return
	}
	preorder(node.left)
	preorder(node.right)
	fmt.Print(node.data + " ")
}
func main() {
	s := "a + b - c"
	st := strings.Split(s, " ")
	root := treeCeator(st)
	preorder(root)
	fmt.Println()
	postorder(root)
	fmt.Println()
}
