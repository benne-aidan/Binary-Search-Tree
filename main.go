package main

import (
	"fmt"
	"math/rand/v2"
	"time"
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

func submission() {
	// Initialize tree
	var head TreeNode
	head.headInit()

	// ================== QUESTION 4: Add values 1-15 to tree and print =====================================
	for i := range 15 {
		head.Insert(int32(i + 1))
	}
	fmt.Print("Tree: ")
	head.PrintTree()
	fmt.Println()

	// ================== QUESTION 5: Track CPU times for Search(1) and Search(15) ==========================
	// Track time for Search(1)
	fmt.Println("Running Search(1) 100,000 times...")
	start := time.Now()
	for range 100000 {
		head.Search(1)
	}
	end := time.Now()

	runtime := end.Sub(start)
	fmt.Println("Search(1) took ", runtime, " for 100,000 iterations")

	// Track time for Search(15)
	fmt.Println("Running Search(15) 100,000 times...")
	start = time.Now()
	for range 100000 {
		head.Search(15)
	}
	end = time.Now()

	runtime = end.Sub(start)
	fmt.Println("Search(15) took ", runtime, " for 100,000 iterations")

	// ================== QUESTION 6: Remove tree keys ======================================================
	// Remove key 5, print
	fmt.Print("Removing 5. New tree: ")
	head.Remove(5)
	head.PrintTree()
	fmt.Println()

	// Remove key 15, print
	fmt.Print("Removing 15. New tree: ")
	head.Remove(15)
	head.PrintTree()
	fmt.Println()

	// Remove key 1, print
	fmt.Print("Removing 1. New tree: ")
	head.Remove(1)
	head.PrintTree()
	fmt.Println()

	// Insert key 2, print
	fmt.Print("Inserting 2. New tree: ")
	head.Insert(2)
	head.PrintTree()
	fmt.Println()

}

func main() {
	//debug()
	submission()

}
