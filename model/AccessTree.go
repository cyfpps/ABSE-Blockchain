package model

// AccessControlTree 访问控制树结构体
type AccessControlTree struct {
	Root *TreeNode // 根节点
}

// TreeNode 访问控制树节点结构体
type TreeNode struct {
	Attribute string      // 节点属性
	Children  []*TreeNode // 子节点列表
}

// NewAccessControlTree 创建一个新的访问控制树
func NewAccessControlTree() *AccessControlTree {
	return &AccessControlTree{
		Root: nil,
	}
}

// AddNode 添加节点到访问控制树
func (act *AccessControlTree) AddNode(parent *TreeNode, attribute string) *TreeNode {
	node := &TreeNode{
		Attribute: attribute,
		Children:  []*TreeNode{},
	}
	if parent == nil {
		act.Root = node
	} else {
		parent.Children = append(parent.Children, node)
	}
	return node
}

// CheckAccess 检查是否满足访问策略
func (act *AccessControlTree) CheckAccess(attributes []string) bool {
	if act.Root == nil {
		return false
	}
	return act.checkAccessRecursive(act.Root, attributes)
}

// 递归地检查访问策略
func (act *AccessControlTree) checkAccessRecursive(node *TreeNode, attributes []string) bool {
	if node == nil {
		return false
	}
	if len(node.Children) == 0 {
		// 叶子节点，检查属性是否匹配
		for _, attr := range attributes {
			if attr == node.Attribute {
				return true
			}
		}
		return false
	}
	// 非叶子节点，递归检查子节点
	for _, child := range node.Children {
		if act.checkAccessRecursive(child, attributes) {
			return true
		}
	}
	return false
}
