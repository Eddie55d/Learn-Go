package main

import (
	"fmt"
)

type Node struct {
	key                 int
	left, right, parent *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(key int) *Node {

	node := &Node{key: key}

	if t.root == nil {
		t.root = node

		return nil
	}

	if t.root != nil {
		if t.root.key > node.key && t.root.left == nil {
			node.parent = t.root
			t.root.left = node
			return nil
		}
		if t.root.key < node.key && t.root.right == nil {
			node.parent = t.root
			t.root.right = node
			return nil
		}

		shouldReturn, returnValue := addNode(t.root, node)
		if shouldReturn {
			return returnValue
		}

	}

	return nil
}

func addNode(current *Node, newNode *Node) (bool, *Node) {

	if newNode.key > current.key {
		if current.right != nil {
			addNode(current.right, newNode)
		} else {
			newNode.parent = current
			current.right = newNode
			return true, current
		}
	}

	if newNode.key < current.key {
		if current.left != nil {
			addNode(current.left, newNode)
		} else {
			newNode.parent = current
			current.left = newNode
			return true, newNode
		}
	}

	return false, nil
}

func (t *Tree) Remove(delNode int) (*Node, error) {
	if t.root == nil {
		return nil, fmt.Errorf("tree is empty")
	}

	node, err := t.Search(t.root, delNode)
	if err != nil {
		return nil, err
	}

	if node != nil {
		parent := node.parent
		//Deleting a tree leaf
		if node.left == nil && node.right == nil {
			if parent.left == node {
				parent.left = nil

			}
			if parent.right == node {
				parent.right = nil

			}

			return node, nil
		} else if node.left == nil || node.right == nil {
			//Deleting a tree node with one child
			if node.left == nil {
				if parent.left == node {
					parent.left = node.right
				} else {
					parent.right = node.right
				}
				node.right.parent = parent
				return node, nil
			} else {
				if parent.left == node {
					parent.left = node.left
				} else {
					parent.right = node.left
				}
				node.left.parent = parent
				return node, nil
			}
		} else {
			//deleting a node with two children
			successor := t.Next(delNode)
			oldNode := &Node{
				key:    node.key,
				left:   node.left,
				right:  node.right,
				parent: node.parent,
			}
			node.key = successor.key
			if successor != nil && successor.parent.left == successor {
				successor.parent.left = successor.right
				if successor.right != nil {
					successor.right.parent = successor.parent
				}
				return oldNode, nil
			} else {
				successor.parent.right = successor.right
				if successor.right != nil {
					successor.right.parent = successor.parent
				}
				return oldNode, nil
			}
		}
	}

	return nil, nil
}

func (n *Node) inorderTraversal() {
	if n == nil {
		return
	}
	n.left.inorderTraversal()
	fmt.Printf("%d\n", n.key)
	n.right.inorderTraversal()

}

func (t *Tree) Search(node *Node, key int) (*Node, error) {

	if t.root == nil {
		return nil, fmt.Errorf("tree is empty")
	}

	if key == node.key && node != nil {
		return node, nil
	}

	if key < node.key && node.left != nil {
		return t.Search(node.left, key)
	} else if key > node.key && node.right != nil {
		return t.Search(node.right, key)
	} else {
		return nil, fmt.Errorf("node not found")
	}

}

func (t *Tree) Next(x int) *Node {
	current := t.root
	successor := new(Node)

	for current != nil {
		if current.key > x {
			successor = current
			current = current.left
		} else {
			current = current.right
		}
	}
	return successor
}

func main() {

	bst := new(Tree)
	fmt.Println(bst.Search(bst.root, 1))

	treeArr := [...]int{8, 3, 10, 1, 6, 4, 7, 10, 14, 13}

	for _, v := range treeArr {
		bst.Insert(v)
	}

	fmt.Println(bst.root.left)
	fmt.Println(bst.root.right)

	bst.root.inorderTraversal()

	fmt.Println(bst.Search(bst.root, 44))

	fmt.Println(bst.Search(bst.root, 14))

	fmt.Println(bst.Remove(14))
	fmt.Println(bst.Search(bst.root, 14))
	fmt.Println(bst.Search(bst.root, 13))

	fmt.Println(bst.Remove(6))
	fmt.Println(bst.Remove(6))

	bst.root.inorderTraversal()

}
