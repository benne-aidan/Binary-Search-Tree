package main

import (
	"fmt"
)

// Node in binary search tree
type TreeNode struct {
	val         int32
	init        uint16 // 0 if node has not been initialized with a value, 1 otherwise
	left, right *TreeNode
}

// ====================== REQUIRED METHODS FOR ASSIGNMENT ===================================================

// Other methods

// InOrder traversal to print tree from head. Does not add newline due to recursive nature of method
func (head *TreeNode) PrintTree() {
	if head.isEmpty() {
		fmt.Println("Tree is empty")
		return
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
		head.init = 1 // Update node init
		head.setVal(newVal)
		return
	}

	// Find node to append to
	targetNode := head.findNode(newVal)

	// Determine whether to append left or right, and do so
	if newVal < targetNode.val {
		targetNode.setLeft(newVal)
	} else if newVal > targetNode.val {
		targetNode.setRight(newVal)
	} else { // If not left or right, element is duplicate insert randomly to left or right
		if targetNode.isExternal() {
			targetNode.setLeft(newVal)
		} else if targetNode.hasRight() {
			targetNode.right.Insert(newVal)
		} else if targetNode.hasLeft() {
			targetNode.left.Insert(newVal)
		}
	}
}

// Searches for value in tree. Returns said value if successful, returns last leaf found otherwise
func (head *TreeNode) Search(target int32) int32 {
	return head.findNode(target).val
}

// Removes node from tree
func (head *TreeNode) Remove(target int32) {
	if head.init == 0 {
		panic("ERROR: Cannot remove from empty tree")
	}

	// Get node matching target value
	targetNode := head.findNode(target)

	// If value is not present in tree, throw error
	if targetNode.val != target {
		panic("ERROR: Value not found in tree")
	}

	// NOTE: Make sure algorithm works for head node as well

	// Case: 0 children (leaf node) (works)
	if targetNode.isExternal() {
		parentNode := targetNode.parent(head)

		// Special case when tree becomes empty after removal
		if targetNode.val == parentNode.val {
			head.init = 0
			return
		}

		// Determine whether target is left or right child, delete child
		if target < parentNode.val { // Target is left child, delete left child
			parentNode.left = nil
		} else { // Target is right child, delete right child
			parentNode.right = nil
		}

		// Mark target node as deleted
		targetNode.init = 0
	} else if targetNode.hasLeft() && targetNode.hasRight() { // Case: 2 children
		// Find successor for target node
		swap := targetNode.right.leftMostChild()
		parentNode := swap.parent(head)
		swapNodes(targetNode, swap)

		// Tree looks like this (removing 10)
		/*
				10
			   /  \
			  5   15
			        \
					...
		*/
		if !targetNode.right.hasLeft() {
			targetNode.right = swap.right
			swap.init = 0
		} else {
			// swap is a left child with no left child
			parentNode.left = swap.right
			swap.init = 0
		}

	} else { // Case: 1 child
		// Get parent node to help with deletion
		parentNode := targetNode.parent(head)

		// Determine successor node (either right or left)
		var successor *TreeNode
		if targetNode.hasLeft() {
			successor = targetNode.left
		} else {
			successor = targetNode.right
		}

		// Determine whether targetNode is left or right child to update parentNode reference
		if targetNode.val < parentNode.val {
			parentNode.left = successor
			targetNode.init = 0
		} else if targetNode.val > parentNode.val {
			parentNode.right = successor
			targetNode.init = 0
		} else {
			// In this case the target node is the parent node, which only
			// happens in the case of removing the root node.
			targetNode.val = successor.val
			targetNode.left = successor.left
			targetNode.right = successor.right
			successor.init = 0
		}
	}

}

// ====================== Methods on tree nodes =============================================================

// Initialize head node (empty value)
func (node *TreeNode) headInit() {
	// We denote an empty tree (or subtree) by a node with height zero
	// This allows us to initialize a tree as a head node with nothing in it
	node.left = nil
	node.right = nil
	node.init = 0
}

// Initialize empty node
func (node *TreeNode) nodeInit(newVal int32) {
	node.left = nil
	node.right = nil
	node.val = newVal
	node.init = 1
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
	return head.init == 0
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

// Return pointer to parent of a node. Takes head of tree as arg
func (node *TreeNode) parent(head *TreeNode) *TreeNode {
	if head.init == 0 {
		panic("ERROR: Tree has not been initialized")
	}

	// Case: `child` has no parent, return self
	if head.val == node.val {
		return head
	}

	// Node is to the left
	if node.val < head.val {
		// Parent has been found
		if head.left.val == node.val {
			return head
		} else { // Not found, recurse to left subtree
			return node.parent(head.left)
		}
	} else { // Node is to the right
		// Parent has been found
		if head.right.val == node.val {
			return head
		} else { // Not found, recurse to right subtree
			return node.parent(head.right)
		}
	}
}

// Deletes external node. Supposes that node has a parent
func (node *TreeNode) deleteExt(parentNode *TreeNode) {
	if node.isEmpty() {
		panic("ERROR: deleteExt called on empty tree")
	}

	// Determine whether node is left or right child
	if node.val < parentNode.val {
		parentNode.left = nil
	} else {
		parentNode.right = nil
	}

	node.init = 0
}

// Returns leftmost descendant of node
func (node *TreeNode) leftMostChild() *TreeNode {
	if node.isEmpty() {
		panic("ERROR: leftMostChild called on empty tree")
	}
	if node == nil {
		panic("ERROR: leftMostChild called on nil node")
	}

	if node.hasLeft() {
		return node.left.leftMostChild()
	} else {
		return node
	}
}

// Returns true if tree contains target, false otherwise
func (head *TreeNode) contains(target int32) bool {
	return head.Search(target) == target
}

// ====================== Other functions relating to trees =================================================

// Swap the values of two nodes
func swapNodes(node1, node2 *TreeNode) {
	temp := node1.val
	node1.val = node2.val
	node2.val = temp
}
