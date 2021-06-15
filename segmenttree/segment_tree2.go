package segmenttree

type Node struct {
	left      int
	right     int
	value     interface{}
	leftNode  *Node
	rightNode *Node
}
