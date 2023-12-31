package avltree

import "fmt"

type AVLTree struct {
    Root *Node
}

type Node struct {
    Value int
    Left *Node
    Right *Node
    Parent *Node
}

func (tree *AVLTree) Insert(value int) {
    if tree.Root == nil {
        tree.Root = &Node{Value: value, Left: nil, Right: nil, Parent: nil}
        return
    }
    
    insertInner(&tree.Root, value, tree, nil)
}

func insertInner(node **Node, value int, tree *AVLTree, parent *Node) {
    if *node == nil {
        *node = &Node{Value: value, Left: nil, Right: nil, Parent: parent}
    } else if value < (*node).Value {
        insertInner(&(*node).Left, value, tree, *node)
    } else if value > (*node).Value {
        insertInner(&(*node).Right, value, tree, *node)
    }

    (*node).Balance(tree)
}

func (node *Node) Balance(tree *AVLTree) {
    if node == nil {
        return
    }
    
    balanceFactor := Height(node.Right) - Height(node.Left)

    if balanceFactor > 1 {
        if Height(node.Right.Right) - Height(node.Right.Left) < 0 {
            rightRotate(node.Right, tree)
            leftRotate(node, tree)
        } else {
            leftRotate(node, tree)
        }
    } else if balanceFactor < -1 {
        if Height(node.Left.Right) - Height(node.Left.Left) > 0 {
            leftRotate(node.Left, tree)
            rightRotate(node, tree)
        } else {
            rightRotate(node, tree)
        }
    }
}

func leftRotate(node *Node, tree *AVLTree) {
    right := node.Right
    
    if node.Parent != nil {
        if node.Parent.Value < node.Value {
            node.Parent.Right = right
        } else if node.Parent.Value > node.Value {
            node.Parent.Left = right
        }
    } else {
        node.Parent = right
        right.Parent = nil
        tree.Root = right
    }
    
    node.Right = right.Left
    right.Left = node

    if node.Right != nil {
        node.Right.Parent = node
    }
}

func rightRotate(node *Node, tree *AVLTree) {
    left := node.Left
    
    if node.Parent != nil {
        if node.Parent != nil && node.Parent.Value < node.Value {
            node.Parent.Right = left
        } else if node.Parent != nil && node.Parent.Value > node.Value {
            node.Parent.Left = left
        }
    } else {
        node.Parent = left
        left.Parent = nil
        tree.Root = left
    }

    node.Left = left.Right
    left.Right = node

    if node.Left != nil {
        node.Left.Parent = node
    }
}

func Height(root *Node) int {
    if root == nil {
        return 0
    }

    return max(Height(root.Left), Height(root.Right)) + 1
}

func printUtil(root *Node, space int) {
    if (root == nil) {
        return
    }
 
    space += 10;
 
    printUtil(root.Right, space);
 
    fmt.Print("\n")

    for  i := 10; i < space; i++ {
        fmt.Print(" ")
    }

    fmt.Printf("%d\n", root.Value);
 
    printUtil(root.Left, space);
}
 
func (tree *AVLTree) Print() {
    printUtil(tree.Root, 0);
}
 
