package MODEL

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// Node 表示访问控制树中的一个节点
type Node struct {
	Attribute  string  // 叶子节点的属性名称
	Threshold  int     // 非叶子节点的阈值
	Children   []*Node // 子节点
	Polynomial []int   // 非叶子节点的多项式系数
}

// ACT 表示访问控制树
type ACT struct {
	Root *Node
}

// NewNode 创建一个新节点
func NewNode(attribute string, threshold int) *Node {
	return &Node{
		Attribute: attribute,
		Threshold: threshold,
		Children:  []*Node{},
	}
}

// AddChild 向节点添加一个子节点
func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

// SetPolynomial 为节点设置多项式
func (n *Node) SetPolynomial() {
	degree := n.Threshold - 1
	n.Polynomial = make([]int, degree+1)

	// 随机选择多项式系数的值
	for i := 0; i <= degree; i++ {
		randomValue, _ := rand.Int(rand.Reader, big.NewInt(100))
		n.Polynomial[i] = int(randomValue.Int64())
	}
}

// EvaluatePolynomial 评估给定x的多项式
func (n *Node) EvaluatePolynomial(x int) int {
	result := 0
	for i, coef := range n.Polynomial {
		result += coef * pow(x, i)
	}
	return result
}

// pow 计算x的y次方
func pow(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}

// ParsePolicy 解析策略字符串并构建访问控制树
func ParsePolicy(policy string) *Node {
	policy = strings.TrimSpace(policy)
	if strings.Contains(policy, "AND") {
		parts := strings.Split(policy, "AND")
		node := NewNode("", len(parts))
		for _, part := range parts {
			node.AddChild(ParsePolicy(part))
		}
		node.SetPolynomial()
		return node
	} else if strings.Contains(policy, "OR") {
		parts := strings.Split(policy, "OR")
		node := NewNode("", 1) // 对于OR，阈值始终为1
		for _, part := range parts {
			node.AddChild(ParsePolicy(part))
		}
		node.SetPolynomial()
		return node
	} else {
		return NewNode(strings.TrimSpace(policy), 0)
	}
}

// Satisfies 检查给定的属性集是否满足节点的访问策略
func Satisfies(node *Node, attributes []string) bool {
	if node.Attribute != "" {
		for _, attr := range attributes {
			if node.Attribute == attr {
				return true
			}
		}
		return false
	} else {
		count := 0
		for _, child := range node.Children {
			if Satisfies(child, attributes) {
				count++
			}
		}
		return count >= node.Threshold
	}
}

// PrintTree 打印访问控制树的结构
func PrintTree(node *Node, level int) {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += "  "
	}
	if node.Attribute != "" {
		fmt.Println(prefix + node.Attribute)
	} else {
		fmt.Println(prefix + fmt.Sprintf("T(%d)", node.Threshold))
		for _, child := range node.Children {
			PrintTree(child, level+1)
		}
	}
}
