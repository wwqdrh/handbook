package tree

import (
	"container/list"
	"fmt"
)

// 红黑树是一种近似平衡的二叉查找树，从2-3树或2-3-4树衍生而来。
// 通过对二叉树节点进行染色，染色为红或黑节点，来模仿2-3树或2-3-4树的3节点和4节点，从而让树的高度减小。
// 2-3-4树对照实现的红黑树是普通的红黑树，而2-3树对照实现的红黑树是一种变种，称为左倾红黑树，其更容易实现。
// 自平衡二叉查找树

// 红黑树是每个节点都带有颜色属性（红色或黑色）的二叉查找树。红黑树也属于自平衡二叉查找树。
// 红黑树具有如下性质：
// 1、每个节点要么是红色要么是黑色。
// 2、树的根结点为黑色节点。
// 3、所有叶子节点都是黑色节点（叶子是NIL节点）。
// 4、每个红色节点必须有两个黑色的子节点（从每个叶子到根的所有路径上不能有两个连续的红色节点）。
// 5、从任意节点到其每个叶子节点的所有简单路径都包含相同数目的黑色节点。

// 普通红黑树: 允许一个节点有两个红色的子节点 对应2-3-4树
// 左倾红黑树: 一个节点只能有一个红色子节点 并且是做节点 对应2-3树

type Entryer interface {
	GetValue() interface{}
	Compare(Entryer) int
}

////////////////////
// node
////////////////////
type Color uint8

const (
	RED Color = iota + 1
	BLACK
)

type RBNode struct {
	entry               Entryer
	color               Color
	parent, left, right *RBNode
}

func NewRBNode(entry Entryer) *RBNode {
	rbNoe := &RBNode{
		entry:  entry,
		color:  RED,
		parent: nil,
		left:   nil,
		right:  nil,
	}
	return rbNoe
}

// getGrandParent() 获取父级节点的父级节点
func (rbNode *RBNode) getGrandParent() *RBNode {
	parent := rbNode.parent
	if parent != nil {
		return parent.parent
	} else {
		return nil
	}
}

func (rbNode *RBNode) GetEntry() Entryer {
	return rbNode.entry
}

// getSibling() 获取兄弟节点
func (rbNode *RBNode) getSibling() *RBNode {
	parent := rbNode.parent
	if parent != nil {
		if rbNode == parent.left {
			return parent.right
		} else {
			return parent.left
		}
	} else {
		return nil
	}
}

// GetUncle() 父节点的兄弟节点
func (rbNode *RBNode) getUncle() *RBNode {
	parent := rbNode.parent
	if parent != nil {
		return parent.getSibling()
	} else {
		return nil
	}
}

//左旋参数为旋转轴的节点 若根节点变动返回根节点
func (rbNode *RBNode) leftRotate() *RBNode {
	var root *RBNode
	if rbNode == nil {
		return root
	}
	if rbNode.right == nil {
		return root
	}
	parent := rbNode.parent
	var isLeft bool
	if parent != nil {
		isLeft = parent.left == rbNode
	}
	grandson := rbNode.right.left
	if rbNode.right.left != nil {
		rbNode.right.left.parent = rbNode
	}
	rbNode.right.left = rbNode
	rbNode.parent = rbNode.right
	rbNode.right = grandson
	// 判断是否换了根节点
	if parent == nil {
		rbNode.parent.parent = nil
		root = rbNode.parent
	} else {
		if isLeft {
			parent.left = rbNode.parent
		} else {
			parent.right = rbNode.parent
		}
		rbNode.parent.parent = parent
	}
	return root
}

//右旋参数为旋转轴的节点 若根节点变动返回根节点
func (rbNode *RBNode) rightRotate() *RBNode {
	var root *RBNode
	if rbNode == nil {
		return root
	}
	if rbNode.left == nil {
		return root
	}
	parent := rbNode.parent
	var isLeft bool
	if parent != nil {
		isLeft = parent.left == rbNode
	}
	grandson := rbNode.left.right
	if grandson != nil {
		grandson.parent = rbNode
	}
	rbNode.left.right = rbNode
	rbNode.parent = rbNode.left
	rbNode.left = grandson
	// 判断是否换了根节点
	if parent == nil {
		rbNode.parent.parent = nil
		root = rbNode.parent
	} else {
		if isLeft {
			parent.left = rbNode.parent
		} else {
			parent.right = rbNode.parent
		}
		rbNode.parent.parent = parent
	}
	return root
}

////////////////////
// tree
////////////////////

type RBTree struct {
	root *RBNode
}

//插入操作
func (rbTree *RBTree) Insert(entry Entryer) {
	if rbTree.root == nil {
		root := NewRBNode(entry)
		rbTree.insertCheck(root)
		return
	}
	rbTree.insertNode(rbTree.root, entry)
}

//查询节点
func (rbTree *RBTree) GetNode(entry Entryer) *RBNode {
	pNode := getNode(rbTree.root, entry)
	if pNode == nil {
		return pNode
	}
	result := new(RBNode)
	result.entry = pNode.entry
	return result
}

func (rbTree *RBTree) GetRoot() {
	fmt.Println(rbTree.root)
}

//删除操作
func (rbTree *RBTree) DeleteNode(entry Entryer) bool {
	query := getNode(rbTree.root, entry)
	if query == nil {
		return false
	}
	if query.left == nil || query.right == nil {
		rbTree.deleteOneNode(query)
	} else {
		//要删除的节点有两个子节点的时候 找到右子节点的最左子节点替换 删除该最左子节点变成删除一个节点的操作
		mostLeft := query.right
		for mostLeft.left != nil {
			mostLeft = mostLeft.left
		}
		query.entry = mostLeft.entry //替换
		rbTree.deleteOneNode(mostLeft)
	}
	return true
}

//中序遍历顺序输出
func (rbTree *RBTree) MidRec() []*RBNode {
	res := []*RBNode{}
	midRec(rbTree.root, &res)
	return res
}

//层序遍历输出
func (rbTree *RBTree) LevelTraversal() {
	l := list.New()
	l.PushBack(rbTree.root)
	levelTraversal(l)
}

//层序遍历
func levelTraversal(l *list.List) {
	e := l.Front()
	l.Remove(e)
	for e != nil {
		v := e.Value
		pNode := v.(*RBNode)
		fmt.Print(pNode.entry.GetValue())
		fmt.Print(" ")
		fmt.Print(pNode.color)
		fmt.Print("  ")
		if pNode.left != nil {
			l.PushBack(pNode.left)
		}
		if pNode.right != nil {
			l.PushBack(pNode.right)
		}
		e = l.Front()
		if e != nil {
			l.Remove(e)
		}
	}
}

//中序遍历
func midRec(pNode *RBNode, data *[]*RBNode) {
	if pNode != nil {
		midRec(pNode.left, data)
		*data = append(*data, pNode)
		// fmt.Print(pNode.entry.GetValue())
		// fmt.Print(" ")
		// fmt.Print(pNode.color)
		// fmt.Print("  ")
		midRec(pNode.right, data)
	}
}

//查询节点
func getNode(pNode *RBNode, entry Entryer) *RBNode {
	if pNode == nil {
		return nil
	}
	res := pNode.entry.Compare(entry)
	if res == 0 {
		return pNode
	} else if res == -1 {
		return getNode(pNode.left, entry)
	} else {
		return getNode(pNode.right, entry)
	}
}

//插入节点
func (rbTree *RBTree) insertNode(pNode *RBNode, entry Entryer) {
	res := pNode.entry.Compare(entry)
	if res != 1 {
		if pNode.left != nil {
			rbTree.insertNode(pNode.left, entry)
		} else {
			temp := NewRBNode(entry)
			temp.parent = pNode
			pNode.left = temp
			rbTree.insertCheck(temp)
		}
	} else {
		if pNode.right != nil {
			rbTree.insertNode(pNode.right, entry)
		} else {
			temp := NewRBNode(entry)
			temp.parent = pNode
			pNode.right = temp
			rbTree.insertCheck(temp)
		}
	}
}

//检查插入
func (rbTree *RBTree) insertCheck(pNode *RBNode) {
	parent := pNode.parent
	if parent == nil {
		pNode.color = BLACK
		rbTree.root = pNode
		return
	}
	//父亲节点为红色则要处理
	if parent.color == RED {
		uncle := pNode.getUncle()
		if uncle != nil && uncle.color == RED {
			uncle.color = BLACK
			parent.color = BLACK
			parent.parent.color = RED
			rbTree.insertCheck(parent.parent)
		} else {
			grandParent := pNode.getGrandParent()
			if grandParent.left == parent {
				if parent.right == pNode {
					parent.leftRotate() //父节点先左旋
				}
				if root := grandParent.rightRotate(); root != nil {
					rbTree.root = root
				}
			} else {
				if parent.left == pNode {
					parent.rightRotate() //父节点先右旋
				}
				if root := grandParent.leftRotate(); root != nil {
					rbTree.root = root
				}
			}
			grandParent.color = RED
			parent.color = BLACK
		}
	}
}

//删除一个节点或没有节点的节点
func (rbTree *RBTree) deleteOneNode(rbNode *RBNode) {
	var child *RBNode
	if rbNode.left == nil {
		child = rbNode.right
	} else {
		child = rbNode.left
	}
	parent := rbNode.parent
	if parent == nil {
		if child == nil {
			rbTree.root = nil
		} else {
			child.color = BLACK
			child.parent = nil
			rbTree.root = child
		}
		rbNode = nil
		return
	}
	if rbNode.color == RED {
		if parent.left == rbNode {
			parent.left = child
		} else {
			parent.right = child
		}
		if child != nil {
			child.parent = parent
		}
		rbNode = nil
		return
	}
	if child == nil {
		child = new(RBNode)
		child.parent = parent
		if parent.left == rbNode {
			parent.left = child
			rbTree.deleteCheck(child)
			parent.left = nil
		} else {
			parent.right = child
			rbTree.deleteCheck(child)
			parent.right = nil
		}
		child = nil
		rbNode = nil
		return
	} else {
		if parent.left == rbNode {
			parent.left = child
		} else {
			parent.right = child
		}
		child.parent = parent
		if child.color == RED {
			child.color = BLACK
			rbNode = nil
			return
		}
		rbTree.deleteCheck(child)
		rbNode = nil
		return
	}

}

//删除检查
func (rbTree *RBTree) deleteCheck(rbNode *RBNode) {
	parent := rbNode.parent
	if parent == nil {
		rbNode.color = BLACK
		rbTree.root = rbNode
		return
	}
	brother := rbNode.getSibling()
	if brother.color == RED {
		if parent.left == brother {
			parent.rightRotate()
		} else {
			parent.leftRotate()
		}
		parent.color = RED
		brother.color = BLACK
		brother = rbNode.getSibling()
		parent = rbNode.parent
	}
	s1Color := BLACK
	s2Color := BLACK
	if brother.left != nil {
		s1Color = brother.left.color
	}
	if brother.right != nil {
		s2Color = brother.right.color
	}
	if s1Color == BLACK && s2Color == BLACK {
		if parent.color == RED {
			parent.color = BLACK
			brother.color = RED
			return
		}
		brother.color = RED
		rbTree.deleteCheck(parent)
		return
	}
	if parent.left == rbNode && s1Color == RED && s2Color == BLACK {
		brother.color = RED
		brother.left.color = BLACK
		brother.rightRotate()
	} else if parent.right == rbNode && s1Color == BLACK && s2Color == RED {
		brother.color = RED
		brother.right.color = BLACK
		brother.leftRotate()
	}
	brother.color = parent.color
	parent.color = BLACK
	if parent.left == rbNode {
		brother.right.color = BLACK
		parent.leftRotate()
	} else {
		brother.left.color = BLACK
		parent.rightRotate()
	}
}
