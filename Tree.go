package main

// Node in binary search tree
type TreeNode[T comparable] struct {
	val         T
	height      uint16
	left, right *TreeNode[T]
}

// ====================== REQUIRED METHODS FOR ASSIGNMENT ===================================================

// Other methods

// ====================== Methods on tree nodes =============================================================

// Initialize head node (empty value)
func (node *TreeNode[T]) headInit() {
	node.left = nil
	node.right = nil
	node.height = 0
}

// Initialize empty node
func (node *TreeNode[T]) nodeInit(newVal T) {
	node.left = nil
	node.right = nil
	node.val = newVal
	node.height = 1
}

// Return true if node has left child
func (node *TreeNode[T]) hasLeft() bool {
	if node.left == nil {
		return false
	} else {
		return true
	}
}

// Return true if node has right child
func (node *TreeNode[T]) hasRight() bool {
	if node.right == nil {
		return false
	} else {
		return true
	}
}

// Return pointer to left child of node
func (node *TreeNode[T]) getLeft() *TreeNode[T] {
	return node.left
}

// Return pointer to right child of node
func (node *TreeNode[T]) getRight() *TreeNode[T] {
	return node.right
}

// Return true if node is external
func (node *TreeNode[T]) isExternal() bool {
	return !(node.hasLeft() || node.hasRight())
}

// Update value stored in node
func (node *TreeNode[T]) setVal(newVal T) {
	node.val = newVal
}

// Set left node to a particular value
func (node *TreeNode[T]) setLeft(newVal T) {
	newNode := &TreeNode[T]{}
	newNode.nodeInit(newVal)

	node.left = newNode
}

// Set right node to a particular value
func (node *TreeNode[T]) setRight(newVal T) {
	newNode := &TreeNode[T]{}
	newNode.nodeInit(newVal)

	node.right = newNode
}

/*
func print() {
	if hasLeft {
		left.print()
	}
	System.out.println(val)
	if hasRight {
		right.print()
	}
}
*/
