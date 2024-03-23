package main

func main() {
	var head TreeNode
	head.headInit()
	head.setVal(5)
	head.setLeft(2)

	head.setRight(7)
	head.right.setLeft(6)
	head.left.setRight(3)

	head.PrintTree()

	//if head.hasLeft() {
	//	fmt.Println("Has child")
	//	fmt.Println(head.left.val)
	//} else {
	//	fmt.Println("Childless")
	//}
}
