package main

import "fmt"

func main() {
	var head TreeNode
	head.headInit()
	head.Insert(5)
	head.Insert(3)
	head.Insert(7)
	head.Insert(2)
	head.Insert(4)
	head.Insert(6)
	head.Insert(8)

	head.PrintTree()
	fmt.Println()

	fmt.Println(head.Search(4))
	fmt.Println(head.Search(1))

	//if head.hasLeft() {
	//	fmt.Println("Has child")
	//	fmt.Println(head.left.val)
	//} else {
	//	fmt.Println("Childless")
	//}
}
