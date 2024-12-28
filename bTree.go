package main

/// BTREE SIMPLE IMPLEMENTATION /**/

/*
BTREE-
	-NODE
		|-BTREE_NODE_STRUCTURE
			|-T_ORDER_OF_BTREE , (each node has atleast t-1 keys and atmost 2*t-1  keys )
			|-keys, (slice of integers to store keys in each node)
			|-children, (slice of pointers to BTREE child nodes)
			|-leaf , A boolean value indicating whether the node is a leaf node or not
		|-BTREE_NODE_METHODS
			|-NewBTREENode ,(CONSTRUCTOR) (function to create a new BTREE node)
			|-InsertNonFull , (function to insert a key in a non-full BTREE node)
			|-SplitCHILD , (function to split a child node and create a new BTREE node)

	-BTREE_STRUCTURE
		|-T_ORDER_OF_BTREE , (order of the BTREE)
		|-root , (pointer to the root node of the BTREE)

	-BTREE_METHODS
		|-NewBTREE ,(CONSTRUCTOR) , (function to create a new BTREE)
		|-Insert , (function to insert a key in a BTREE)
		|-PrintBTREE , (function to print the BTREE) //// FOR VIEW PURPOSE
*/

///NODE

type BtreeNode struct {
	t        int
	keys     []int
	children []*BtreeNode
	isLeaf   bool
}

func (n *BtreeNode) NewBTREENode(t int, isLeaf bool) *BtreeNode {
	return &BtreeNode{
		t:        t,
		keys:     []int{},
		children: []*BtreeNode{},
		isLeaf:   isLeaf,
	}
}
func (n *BtreeNode) InsertNonFull(key int) {

	i := len(n.keys) - 1 /// INSERT IN SORTED POSITION O(N)

	if n.isLeaf { ///if leafnode
		n.keys = append(n.keys, 0) /// expand node keys slice for 1 more entry
		///// SEARCH // replace with binary search
		for i >= 0 && key < n.keys[i] { //instert in sorted positon
			n.keys[i+1] = n.keys[i]
			i--
		}
		n.keys[i+1] = key
	} else { // find appropriate child node to insert key
		///// SEARCH // replace with binary search
		for i >= 0 && key < n.keys[i] {
			i--
		}
		i += 1
		///what to do if length of child node is full?
		//split child and
		//again check for appropriate child node to insert key
	}
	n.children[i].InsertNonFull(key)

}

func (n *BtreeNode) SplitCHILD(i int) *BtreeNode {
	return nil
}
