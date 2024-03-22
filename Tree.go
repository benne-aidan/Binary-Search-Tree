package main

type BST[T comparable] struct {
	head *element[T]
}

type element[T comparable] struct {
	val         T
	left, right *element[T]
}

// ====================== REQUIRED METHODS FOR ASSIGNMENT ===================================================
// Methods go here

// Other methods

// ====================== Methods on BSTs ===================================================================

// Initialize empty tree
func (tree *BST[T]) treeInit() {
	tree.head = &element[T]{}
}

// ====================== Methods on tree nodes =============================================================

// Initialize empty node
func (node *element[T]) nodeInit(newVal T) {
	node.left = nil
	node.right = nil
	node.val = newVal
}

// Return true if node has left child
func (node *element[T]) hasLeft() bool {
	if node.left == nil {
		return false
	} else {
		return true
	}
}

// Return true if node has right child
func (node *element[T]) hasRight() bool {
	if node.right == nil {
		return false
	} else {
		return true
	}
}

// Return pointer to left child of node
func (node *element[T]) getLeft() *element[T] {
	return node.left
}

// Return pointer to right child of node
func (node *element[T]) getRight() *element[T] {
	return node.right
}

// Return true if node is external
func (node *element[T]) isExternal() bool {
	return !(node.hasLeft() || node.hasRight())
}

// Update value stored in node
func (node *element[T]) setVal(newVal T) {
	node.val = newVal
}

// Set left node to a particular value
func (node *element[T]) setLeft(newVal T) {
	newNode := &element[T]{}
	newNode.nodeInit(newVal)

	node.left = newNode
}

// Set right node to a particular value
func (node *element[T]) setRight(newVal T) {
	newNode := &element[T]{}
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
