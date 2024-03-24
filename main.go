package main

import (
	"fmt"
	"math/rand/v2"
)

func debug() {
	var head TreeNode
	head.headInit()

	// General test case to test tree as whole. Uncomment to test random insertions and deletions
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

func main() {
	debug()
	// TODO: Define function for test cases listed in assignment, measure runtime

}
