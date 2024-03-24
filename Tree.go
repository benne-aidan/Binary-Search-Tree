package main

import "fmt"

// Node in binary search tree
type TreeNode struct {
	val         int32
	height      uint16
	left, right *TreeNode
}

// ====================== REQUIRED METHODS FOR ASSIGNMENT ===================================================

// Other methods

// InOrder traversal to print tree from head. Does not add newline due to recursive nature of method
func (head *TreeNode) PrintTree() {
	if head.isEmpty() {
		fmt.Println("Tree is empty")
	}

	if head.hasLeft() {
		head.left.PrintTree()
	}
	fmt.Print(head.val)
	fmt.Print(" ")
	if head.hasRight() {
		head.right.PrintTree()
	}
}

// Insert node into tree from head. Ignores duplicate elements
func (head *TreeNode) Insert(newVal int32) {
	// Special case for empty tree
	if head.isEmpty() {
		head.height++ // Update node height
		head.setVal(newVal)
		return
	}

	head.height++ // Update node height

	// Find node to append to
	targetNode := head.findNode(newVal)

	// Determine whether to append left or right, and do so
	if newVal < targetNode.val {
		targetNode.setLeft(newVal)
	} else if newVal > targetNode.val {
		targetNode.setRight(newVal)
	}
}

// Searches for value in tree. Returns said value if successful, returns last leaf found otherwise
func (head *TreeNode) Search(target int32) int32 {
	return head.findNode(target).val
}

// ====================== Methods on tree nodes =============================================================

// Initialize head node (empty value)
func (node *TreeNode) headInit() {
	// We denote an empty tree (or subtree) by a node with height zero
	// This allows us to initialize a tree as a head node with nothing in it
	node.left = nil
	node.right = nil
	node.height = 0
}

// Initialize empty node
func (node *TreeNode) nodeInit(newVal int32) {
	node.left = nil
	node.right = nil
	node.val = newVal
	node.height = 1
}

// Return true if node has left child
func (node *TreeNode) hasLeft() bool {
	if node.left == nil {
		return false
	} else {
		return true
	}
}

// Return true if node has right child
func (node *TreeNode) hasRight() bool {
	if node.right == nil {
		return false
	} else {
		return true
	}
}

// Return true if node is external
func (node *TreeNode) isExternal() bool {
	return !(node.hasLeft() || node.hasRight())
}

// Update value stored in node
func (node *TreeNode) setVal(newVal int32) {
	node.val = newVal
}

// Set left node to a particular value
func (node *TreeNode) setLeft(newVal int32) {
	newNode := &TreeNode{}
	newNode.nodeInit(newVal)

	node.left = newNode
}

// Set right node to a particular value
func (node *TreeNode) setRight(newVal int32) {
	newNode := &TreeNode{}
	newNode.nodeInit(newVal)

	node.right = newNode
}

// Returns true if tree is empty, false otherwise
func (head *TreeNode) isEmpty() bool {
	return head.height == 0
}

// Searches for node in tree containing targetVal. Returns node if successful, returns last node found otherwise
func (head *TreeNode) findNode(targetVal int32) *TreeNode {
	// Throw error if tree is empty
	if head.isEmpty() {
		panic("ERROR: Cannot search empty tree")
	}

	// Base case: node is external (target not found)
	if head.isExternal() {
		return head
	}

	// If target is less than node
	if targetVal < head.val {
		// If no value less than node is present in tree
		if !head.hasLeft() {
			return head
		}
		// Otherwise, recurse and search left subtree
		return head.left.findNode(targetVal)

	} else if targetVal > head.val { // Target is greater than node
		// If no value greater than node is present in tree
		if !head.hasRight() {
			return head
		}
		// Otherwise, recurse and search right subtree
		return head.right.findNode(targetVal)
	} else { // Base case: node is found
		return head
	}
}
