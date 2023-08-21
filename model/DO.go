package model

import (
	"github.com/Nik-U/pbc"
)

// DataOwner 数据拥有者结构体
type DataOwner struct {
	PrivateKey *pbc.Element
	PKDO       *pbc.Element
}

/*
// NewDataOwner 创建一个数据拥有者实例
func NewDataOwner(g *pbc.Element) (*DataOwner, error) {

	/*privateKey := GenerateRandomElement()
	fmt.Println("撤销失123。")
	//	pkDO := new(pbc.Element).PowZn(g, privateKey)
	// 将 privateKey 扩大三倍
	expandedPrivateKey := new(pbc.Element).PowZn(privateKey, big.NewInt(3))

	// 计算 pkDO，使用扩大后的 privateKey 进行指数运算
	pkDO := new(pbc.Element).PowZn(g, expandedPrivateKey)

	fmt.Println("撤销失败11。")
	return &DataOwner{
		PrivateKey: privateKey,
		PKDO:       pkDO,
	}, nil
}
*/
