package tree

type BinarySearchTreeNode struct {
	Data  interface{}
	Left  *BinarySearchTreeNode
	Right *BinarySearchTreeNode
}

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}

func NewBinarySearchTree(DataArr ...interface{}) *BinarySearchTree {
	t := &BinarySearchTree{}
	for _, data := range DataArr {
		t.Insert(data)
	}
	return t
}

func (t *BinarySearchTree) Insert(Data interface{}) {

	if t.Root == nil {
		t.Root = &BinarySearchTreeNode{
			Data:  Data,
			Left:  nil,
			Right: nil,
		}

		return
	}

	tNode := t.Root

	for {
		if tNode.Data == Data {
			return
		}
		if Data.(int) > tNode.Data.(int) {
			if tNode.Right == nil {
				tNode.Right = &BinarySearchTreeNode{
					Data:  Data,
					Left:  nil,
					Right: nil,
				}
				return
			}
			tNode = tNode.Right
		} else {
			if tNode.Left == nil {
				tNode.Left = &BinarySearchTreeNode{
					Data:  Data,
					Left:  nil,
					Right: nil,
				}
				return
			}
			tNode = tNode.Left
		}
	}
}

func (t *BinarySearchTree) Delete(Data interface{}) {
	var pnode *BinarySearchTreeNode
	node := t.Root

	for {
		if node == nil {
			return
		}
		if node.Data == Data {
			break
		} else if node.Data.(int) < Data.(int) {
			pnode = node
			node = node.Right
		} else {
			pnode = node
			node = node.Left
		}
	}

	if node.Left == nil && node.Right == nil {
		if pnode == nil {
			t.Root = nil
		} else if pnode.Data.(int) > node.Data.(int) {
			pnode.Left = nil
		} else {
			pnode.Right = nil
		}
	} else if node.Left != nil && node.Right == nil {
		if pnode == nil {
			t.Root = node.Left
		} else if pnode.Data.(int) > node.Data.(int) {
			pnode.Left = node.Left
		} else {
			pnode.Right = node.Left
		}
	} else if node.Left == nil && node.Right != nil {
		if pnode == nil {
			t.Root = node.Right
		} else if pnode.Data.(int) > node.Data.(int) {
			pnode.Left = node.Right
		} else {
			pnode.Right = node.Right
		}
	} else {
		//右子树最小节点，最小节点不可能存在左子树
		minNode := node.Right
		//右子树最小节点的父节点
		minPNode := node
		for minNode.Left != nil {
			minPNode = minNode
			minNode = minNode.Left
		}
		//最小节点的父节点的左节点指向最小节点的右子树，
		if minPNode.Data.(int) < minNode.Data.(int) {
			minPNode.Right = minNode.Right
		} else {
			minPNode.Left = minPNode.Right
		}

		node.Data = minNode.Data
	}
}
