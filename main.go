package main

import "fmt"

func main() {
	var head TreeNode[int32]
	head.headInit()
	head.setVal(5)
	head.setLeft(2)

	if head.hasLeft() {
		fmt.Println("Has child")
		fmt.Println(head.left.val)
	} else {
		fmt.Println("Childless")
	}
}
