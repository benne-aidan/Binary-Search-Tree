package main

import "fmt"

func main() {
	var list BST[int32]
	list.treeInit()
	list.head.setVal(5)
	list.head.setLeft(2)

	if list.head.hasLeft() {
		fmt.Println("Has child")
		fmt.Println(list.head.left.val)
	} else {
		fmt.Println("Childless")
	}
}
