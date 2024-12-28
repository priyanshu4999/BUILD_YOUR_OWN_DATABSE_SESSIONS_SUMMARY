# B-tree Discussion Summary

## 1. Overview of B-trees and Operations
- **B-trees**: Height-balanced, all leaf nodes at the same height.
- **Height-Balanced Trees**: AVL and RB trees are height-balanced with height \(O(\log N)\), making lookups \(O(\log N)\).
- **Generalizing Binary Trees**: n-ary trees like the 2-3-4 tree are extensions of binary trees, equivalent to RB trees.

## 2. B+ Trees and Nested Arrays
- **B+ Tree Structure**: Internal nodes contain keys indicating key ranges, leaf nodes contain values.
- **Nested Arrays**: Splitting a sorted array into smaller arrays reduces update complexity from \(O(N)\) to \(O(N/m)\). Using B+ trees for nested arrays maintains efficiency.

## 3. Maintaining B+ Trees
- **Invariants**:
  1. Same height for all leaf nodes.
  2. Node size bounded by a constant.
  3. Nodes are not empty.
- **Growing by Splitting Nodes**: Inserting can cause splits, which may propagate up, increasing tree height.
- **Shrinking by Merging Nodes**: Deleting can create empty nodes, restored by merging with siblings to maintain balance.

## 4. Operations on Arrays and Their Time Complexity
- **Access by Index**: \(O(1)\)
- **Search**: 
  - Unsorted: \(O(N)\)
  - Sorted (Binary Search): \(O(\log N)\)
- **Insertion**:
  - At End: \(O(1)\)
  - At Beginning/Middle: \(O(N)\)
- **Deletion**:
  - From End: \(O(1)\)
  - From Beginning/Middle: \(O(N)\)
- **Updates**:
  - By Index: \(O(1)\)
  - By Value (Unsorted): \(O(N)\)
  - By Value (Sorted): \(O(\log N)\)

## 5. Code Example: B-tree in Go
### Components:
- **BTreeNode Struct**:
  - Represents a node in a B-tree.
  - Methods: `NewBTreeNode`, `InsertNonFull`, `SplitChild`.
- **BTree Struct**:
  - Represents the entire B-tree.
  - Methods: `NewBTree`, `Insert`, `PrintTree`.

### Code:
```go
package main

import "fmt"

// BTreeNode represents a node in a B-tree.
type BTreeNode struct {
    t        int
    keys     []int
    children []*BTreeNode
    leaf     bool
}

// NewBTreeNode creates a new BTreeNode.
func NewBTreeNode(t int, leaf bool) *BTreeNode {
    return &BTreeNode{
        t:        t,
        keys:     []int{},
        children: []*BTreeNode{},
        leaf:     leaf,
    }
}

// InsertNonFull inserts a key into a non-full node.
func (n *BTreeNode) InsertNonFull(key int) {
    i := len(n.keys) - 1

    if n.leaf {
        n.keys = append(n.keys, 0)
        for i >= 0 && key < n.keys[i] {
            n.keys[i+1] = n.keys[i]
            i--
        }
        n.keys[i+1] = key
    } else {
        for i >= 0 && key < n.keys[i] {
            i--
        }
        i++
        if len(n.children[i].keys) == 2*n.t-1 {
            n.splitChild(i)
            if key > n.keys[i] {
                i++
            }
        }
        n.children[i].InsertNonFull(key)
    }
}

// SplitChild splits a child node into two nodes.
func (n *BTreeNode) SplitChild(i int) {
    t := n.t
    y := n.children[i]
    z := NewBTreeNode(t, y.leaf)

    n.children = append(n.children[:i+1], append([]*BTreeNode{z}, n.children[i+1:]...)...)
    n.keys = append(n.keys[:i], append([]int{y.keys[t-1]}, n.keys[i:]...)...)

    z.keys = append(z.keys, y.keys[t:]...)
    y.keys = y.keys[:t-1]

    if !y.leaf {
        z.children = append(z.children, y.children[t:]...)
        y.children = y.children[:t]
    }
}

// BTree represents a B-tree.
type BTree struct {
    root *BTreeNode
    t    int
}

// NewBTree creates a new BTree.
func NewBTree(t int) *BTree {
    return &BTree{
        root: NewBTreeNode(t, true),
        t:    t,
    }
}

// Insert inserts a key into the B-tree.
func (t *BTree) Insert(key int) {
    root := t.root
    if len(root.keys) == 2*t.t-1 {
        newRoot := NewBTreeNode(t.t, false)
        newRoot.children = append(newRoot.children, root)
        newRoot.splitChild(0)
        t.root = newRoot
        newRoot.InsertNonFull(key)
    } else {
        root.InsertNonFull(key)
    }
}

// PrintTree prints the B-tree nodes.
func (t *BTree) PrintTree() {
    printNode(t.root, 0)
}

func printNode(node *BTreeNode, level int) {
    fmt.Printf("Level %d: %v\n", level, node.keys)
    for _, child := range node.children {
        printNode(child, level+1)
    }
}

// Main function to demonstrate BTree usage.
func main() {
    btree := NewBTree(3)
    for i := 0; i < 10; i++ {
        btree.Insert(i)
    }
    btree.PrintTree()
}
