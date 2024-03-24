package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var head TreeNode
	head.headInit()

	// Works when root has no children
	// Does it happen with 1 child? Yes

	//	head.Insert(5)
	//	head.Insert(4)
	//	head.Insert(3)
	//	fmt.Print("List: ")
	//	head.PrintTree()
	//	fmt.Println()
	//	head.Remove(5)
	//	fmt.Print("Removed 5. New list: ")
	//	head.PrintTree()
	//	fmt.Println()

	// General test case to test tree as whole. Uncomment when you think you have working code
	for range 10 {
		num := rand.Int32N(50)
		fmt.Println("Inserting ", num)
		head.Insert(num)
	}

	head.PrintTree()
	fmt.Println()

	for i := range 50 {
		if head.isEmpty() {
			break
		}

		if head.contains(int32(i)) {
			head.Remove(int32(i))
			fmt.Printf("Removed %d from tree\nNew tree: ", i)
			head.PrintTree()
			fmt.Println()
		}

	}

}
